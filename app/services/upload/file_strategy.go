package upload

import (
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/startup-of-zero-reais/zoo-api/app/models"
	"golang.org/x/text/unicode/norm"
)

type FileStrategy interface {
	Read() ([]string, error)
	StartImport(models.ImportStatus) error
}

func NewStrategy(reader *csv.Reader) (FileStrategy, error) {
	fileHeader, err := reader.Read()
	if err != nil {
		return nil, err
	}

	if strings.EqualFold(fileHeader[0], "identificação") {
		return &enclosuresStrategy{reader: reader}, nil
	}

	if strings.EqualFold(fileHeader[0], "nome comum") {
		return &speciesStrategy{reader: reader}, nil
	}

	if strings.EqualFold(fileHeader[0], "nome do animal") {
		return &animalsStrategy{reader: reader}, nil
	}

	return nil, fmt.Errorf("can not identify import strategy of this file")
}

func normalize(text string) string {
	text = norm.NFC.String(text)
	text = strings.TrimSpace(text)

	return text
}
