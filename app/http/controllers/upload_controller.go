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
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/services/upload"
)

type UploadController struct {
	UploadService upload.Upload
}

func NewUploadController() *UploadController {
	return &UploadController{
		UploadService: upload.NewUploadService(),
	}
}

func (r *UploadController) Upload(ctx http.Context) http.Response {
	var cf requests.CreateFile

	err := ctx.Request().Bind(&cf)
	if err != nil {
		return ctx.Response().Json(400, http.Json{"error": err.Error()})
	}

	file, err := cf.File.Open()
	if err != nil {
		return ctx.Response().Json(500, http.Json{
			"error": "Failed to open file chunk",
		})
	}

	defer file.Close()

	fileName := fmt.Sprintf("tmp-%s-%d.tmp", cf.UploadID, cf.ChunkIdx)
	filePath := filepath.Join("./tmp/uploads", fileName)

	tmpFile, err := os.Create(facades.App().StoragePath("app/" + filePath))
	if err != nil {
		return ctx.Response().Json(500, http.Json{"error": err.Error()})
	}

	defer tmpFile.Close()

	_, err = io.Copy(tmpFile, file)
	if err != nil {
		return ctx.Response().Json(500, http.Json{"error": err.Error()})
	}

	if cf.ChunkIdx+1 == cf.TotalChunks {
		upDir := facades.App().StoragePath("app/tmp/uploads")
		ffPath := filepath.Join(upDir, fmt.Sprintf("final-%s-%s", cf.UploadID, cf.File.Filename))

		files, err := os.ReadDir(upDir)
		if err != nil {
			return ctx.Response().Json(500, http.Json{"error": err.Error()})
		}

		chunks := sortChunks(files, cf.UploadID)

		ffBuffer := bytes.NewBuffer(nil)
		for _, file := range chunks {
			if fileName != file.Name {
				continue
			}

			chunkPath := filepath.Join(upDir, file.Name)

			cfile, err := os.Open(chunkPath)
			if err != nil {
				return ctx.Response().Json(500, http.Json{"error": err.Error()})
			}

			ffBuffer.ReadFrom(cfile)
			cfile.Close()
		}

		ff, err := os.Create(ffPath)
		if err != nil {
			return ctx.Response().Json(500, http.Json{"error": err.Error()})
		}

		ff.ReadFrom(ffBuffer)
		ff.Close()

		for _, file := range chunks {
			if fileName != file.Name {
				continue
			}

			chunkPath := filepath.Join(upDir, file.Name)

			os.Remove(chunkPath)
		}

		go r.UploadService.Process(ffPath)

	}

	return ctx.Response().Json(200, http.Json{"message": "File uploaded successfully"})
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
			continue
		}

		matches := re.FindStringSubmatch(chunk.Name())
		if matches == nil {
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
