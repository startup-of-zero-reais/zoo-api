package controllers

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/middleware/utils"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
	"github.com/startup-of-zero-reais/zoo-api/app/services/animal"
	"github.com/startup-of-zero-reais/zoo-api/app/services/enclosure"
	"github.com/startup-of-zero-reais/zoo-api/app/services/upload"
)

type UploadController struct {
	UploadService    upload.Upload
	AnimalService    animal.Animal
	EnclosureService enclosure.Enclosure
}

func NewUploadController() *UploadController {
	return &UploadController{
		UploadService:    upload.NewUploadService(),
		AnimalService:    animal.NewAnimalService(),
		EnclosureService: enclosure.NewEnclosureService(),
	}
}

func (r *UploadController) Upload(ctx http.Context) http.Response {
	var cf requests.CreateFile

	err := ctx.Request().Bind(&cf)

	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{"error": err.Error()})
	}

	file, err := cf.File.Open()
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Failed to open file chunk",
		})
	}

	defer file.Close()

	fileName := fmt.Sprintf("tmp-%s-%d.tmp", cf.UploadID, cf.ChunkIdx)
	filePath := filepath.Join("./tmp/uploads", fileName)

	tmpFile, err := os.Create(facades.App().StoragePath("app/" + filePath))
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{"error": err.Error()})
	}

	defer tmpFile.Close()

	_, err = io.Copy(tmpFile, file)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{"error": err.Error()})
	}

	query := facades.Orm().Query()
	importStatus := models.ImportStatus{
		Filename: cf.Filename,
		State:    "sending",
		UploadID: cf.UploadID,
	}

	query.
		Where("upload_id = ?", cf.UploadID).
		FirstOr(&importStatus, func() error {
			return query.Create(&importStatus)
		})

	fmt.Printf("IMPORT STATUS %+v\n", importStatus)

	if cf.ChunkIdx+1 == cf.TotalChunks {
		importStatus.State = "received"
		_, err = query.
			Where("upload_id = ?", cf.UploadID).
			Update(&importStatus)
		if err != nil {
			return ctx.Response().Json(http.StatusInternalServerError, http.Json{"error": err.Error()})
		}

		upDir := facades.App().StoragePath("app/tmp/uploads")
		ffPath := filepath.Join(upDir, fmt.Sprintf("final-%s-%s", cf.UploadID, cf.File.Filename))

		files, err := os.ReadDir(upDir)
		if err != nil {
			return ctx.Response().Json(http.StatusInternalServerError, http.Json{"error": err.Error()})
		}

		chunks := sortChunks(files, cf.UploadID)

		ffBuffer := bytes.NewBuffer(nil)
		for _, file := range chunks {
			if !strings.Contains(file.Name, cf.UploadID) {
				continue
			}

			chunkPath := filepath.Join(upDir, file.Name)

			cfile, err := os.Open(chunkPath)
			if err != nil {
				return ctx.Response().Json(http.StatusInternalServerError, http.Json{"error": err.Error()})
			}

			n, err := ffBuffer.ReadFrom(cfile)
			if err != nil {
				return ctx.Response().Json(http.StatusInternalServerError, http.Json{"error": err.Error()})
			}

			fmt.Println(ffBuffer.String(), n)
			cfile.Close()
		}

		ff, err := os.Create(ffPath)
		if err != nil {
			return ctx.Response().Json(http.StatusInternalServerError, http.Json{"error": err.Error()})
		}

		ff.ReadFrom(ffBuffer)
		ff.Close()

		for _, file := range chunks {
			if !strings.Contains(file.Name, cf.UploadID) {
				continue
			}

			chunkPath := filepath.Join(upDir, file.Name)

			os.Remove(chunkPath)
		}

		go r.UploadService.Process(ffPath, cf, importStatus)
	}

	return ctx.Response().Json(http.StatusOK, http.Json{"message": "File uploaded successfully"})
}

func (r *UploadController) Index(ctx http.Context) http.Response {
	imports, err := r.UploadService.GetImportsStatus()
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{"error": err.Error()})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"total":   len(imports),
		"imports": imports,
	})
}

func (r *UploadController) IndexFiles(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	enclosures, species, animals, err := r.UploadService.GetImportFiles(id)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{"error": err.Error()})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"enclosures": enclosures,
		"species":    species,
		"animals":    animals,
	})
}

func (r *UploadController) ConfirmImport(ctx http.Context) http.Response {
	var cr requests.ConfirmUpload

	aborted := utils.BindAndValidate(ctx, &cr)
	if aborted {
		return nil
	}

	if cr.Type == "animal" {
		return r.processAnimal(ctx, cr.IDs)
	}

	if cr.Type == "enclosure" {
		return r.processEnclosure(ctx, cr.IDs)
	}

	return ctx.Response().Json(http.StatusOK, http.Json{"message": "Confirm uploaded successfully"})
}

type FileEntry struct {
	Name  string
	Index int
}

func sortChunks(chunks []os.DirEntry, uploadID string) []FileEntry {
	var fileEntries []FileEntry

	re := regexp.MustCompile(`tmp-[a-zA-Z0-9]+-(\d+)\.tmp`)

	for _, chunk := range chunks {
		if !strings.Contains(chunk.Name(), uploadID) {
			facades.Log().Errorf("chunk name does not match", chunk.Name(), uploadID)
			continue
		}

		matches := re.FindStringSubmatch(chunk.Name())
		if matches == nil {
			facades.Log().Errorf("chunk name does not matchs", matches)
			continue
		}

		idx, err := strconv.Atoi(matches[1])
		if err != nil {
			facades.Log().Errorf("error on sort chunks %v", err)
			continue
		}

		fileEntry := FileEntry{
			Name:  chunk.Name(),
			Index: idx,
		}

		fileEntries = append(fileEntries, fileEntry)
	}

	sort.Slice(fileEntries, func(i, j int) bool {
		return fileEntries[i].Index < fileEntries[j].Index
	})

	return fileEntries
}

func (r *UploadController) processAnimal(ctx http.Context, ids []string) http.Response {
	animals, err := r.UploadService.GetImportAnimals(ids)
	if err != nil {
		facades.Log().Errorf("failed to import array of animals: %v", err)
		return ctx.Response().Json(http.StatusBadRequest, http.Json{"error": err.Error()})
	}

	err = r.AnimalService.CreateListAnimal(animals)
	if err != nil {
		facades.Log().Errorf("failed to create list animals: %v", err)
		return ctx.Response().Json(http.StatusBadRequest, http.Json{"error": err.Error()})
	}

	err = r.UploadService.RemoveAnimals(ids)

	if err != nil {
		facades.Log().Errorf("failed to deleted import animals: %v", err)
		return ctx.Response().Json(http.StatusBadRequest, http.Json{"error": err.Error()})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{"message": "Import success by model animal"})
}

func (r *UploadController) processEnclosure(ctx http.Context, ids []string) http.Response {
	enclosures, err := r.UploadService.GetImportEnclosures(ids)
	if err != nil {
		facades.Log().Errorf("failed to import array of animals: %v", err)
		return ctx.Response().Json(http.StatusBadRequest, http.Json{"error": err.Error()})
	}

	err = r.EnclosureService.CreateListEnclosure(enclosures)

	if err != nil {
		facades.Log().Errorf("failed to create list enclosures: %v", err)
		return ctx.Response().Json(http.StatusBadRequest, http.Json{"error": err.Error()})
	}

	err = r.UploadService.RemoveEnclosures(ids)

	if err != nil {
		facades.Log().Errorf("failed to deleted import enclosures: %v", err)
		return ctx.Response().Json(http.StatusBadRequest, http.Json{"error": err.Error()})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{"message": "Import success by model enclosure"})
}
