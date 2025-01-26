package upload

import (
	"encoding/csv"
	"os"

	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (e uploadImpl) Process(filePath string, cf requests.CreateFile) {
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

	_, err = query.
		Where("upload_id = ?", cf.UploadID).
		Update(&models.ImportStatus{State: "processing"})
	if err != nil {
		facades.Log().Error(err)
		return
	}

	if err := processor.StartImport(); err != nil {
		facades.Log().Error(err)
		return
	}

	if err := os.Remove(filePath); err != nil {
		facades.Log().Error(err)
		return
	}

	_, err = query.
		Where("upload_id = ?", cf.UploadID).
		Update(&models.ImportStatus{State: "completed"})
	if err != nil {
		facades.Log().Error(err)
	}
}
