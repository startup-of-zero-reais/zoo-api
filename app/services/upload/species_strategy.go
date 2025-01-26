package upload

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

type speciesStrategy struct {
	reader *csv.Reader
}

var _ (FileStrategy) = (*speciesStrategy)(nil)

// Read implements FileStrategy.
func (s *speciesStrategy) Read() ([]string, error) {
	return s.reader.Read()
}

// StartImport implements FileStrategy.
func (s *speciesStrategy) StartImport() error {
	rowCounter := 1 // start at 1 because csv has header line
	doneCounter := 0

	query := facades.Orm().Query()

	for {
		row, err := s.reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			facades.Log().Errorf("failed to read row %v", err)
			return err
		}

		rowCounter = rowCounter + 1

		var species models.ImportSpecies
		species.CommonName = normalize(row[0])
		species.ScientificName = normalize(row[1])
		species.Order = normalize(row[2])
		species.Kind = normalize(row[3])

		if species.CommonName == "" {
			species.Reason = fmt.Sprintf("A espécie da linha %d não está preenchida", rowCounter)
		}

		if species.ScientificName == "" {
			species.Reason = fmt.Sprintf("A espécie da linha %d não está preenchida", rowCounter)
		}

		if species.Order == "" {
			species.Reason = fmt.Sprintf("A espécie da linha %d não está preenchida", rowCounter)
		}

		if species.Kind == "" {
			species.Reason = fmt.Sprintf("A espécie da linha %d não está preenchida", rowCounter)
		}

		err = query.Create(&species)
		if err != nil {
			continue
		}

		doneCounter = doneCounter + 1
	}

	facades.Log().Infof("Imported %d species\n", doneCounter)
	return nil
}
