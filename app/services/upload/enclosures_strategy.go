package upload

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

type enclosuresStrategy struct {
	reader *csv.Reader
}

var _ (FileStrategy) = (*enclosuresStrategy)(nil)

// Read implements FileStrategy.
func (e *enclosuresStrategy) Read() ([]string, error) {
	return e.reader.Read()
}

func (e *enclosuresStrategy) StartImport() error {
	rowCounter := 1 // starts at 1 because csv has header line
	doneCounter := 0

	query := facades.Orm().Query()

	for {
		row, err := e.reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			facades.Log().Errorf("failed to read row %v", err)
			return err
		}

		rowCounter = rowCounter + 1

		var enclosure models.ImportEnclosure
		enclosure.Identification = normalize(row[0])

		if enclosure.Identification == "" {
			enclosure.Reason = fmt.Sprintf("O recinto na linha %d, está vazio, por isso não foi possível importar", rowCounter)
		}

		err = query.Create(&enclosure)
		if err != nil {
			continue
		}

		doneCounter = doneCounter + 1
	}

	facades.Log().Infof("Imported %d enclosures\n", doneCounter)
	return nil
}
