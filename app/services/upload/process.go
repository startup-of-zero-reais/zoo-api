package upload

import (
	"encoding/csv"
	"os"

	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (e uploadImpl) Process(filePath string, cf requests.CreateFile, is models.ImportStatus) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		facades.Log().Errorf("failed to upload process %v", err)
		return
	}

	csvReader := csv.NewReader(file)

	processor, err := NewStrategy(csvReader)
	if err != nil {
		facades.Log().Error(err)
		return
	}

	query := facades.Orm().Query()

	is.State = "processing"
	_, err = query.
		Where("upload_id = ?", cf.UploadID).
		Update(&is)
	if err != nil {
		facades.Log().Error(err)
		return
	}

	if err := processor.StartImport(is); err != nil {
		facades.Log().Error(err)
		return
	}

	if err := os.Remove(filePath); err != nil {
		facades.Log().Error(err)
		return
	}

	is.State = "completed"
	_, err = query.
		Where("upload_id = ?", cf.UploadID).
		Update(&is)
	if err != nil {
		facades.Log().Error(err)
	}
}
