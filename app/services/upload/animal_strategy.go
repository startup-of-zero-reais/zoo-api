package upload

import (
	"encoding/csv"
	"io"
	"time"

	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/helpers"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

type animalsStrategy struct {
	reader *csv.Reader
}

var _ (FileStrategy) = (*animalsStrategy)(nil)

// Read implements FileStrategy.
func (a *animalsStrategy) Read() ([]string, error) {
	return a.reader.Read()
}

// StartImport implements FileStrategy.
func (a *animalsStrategy) StartImport(is models.ImportStatus) error {
	rowCounter := 1 // starts at 1 because csv has header line
	doneCounter := 0

	query := facades.Orm().Query()

	enclosures, err := getEnclosuses()
	if err != nil {
		return err
	}

	species, err := getSpecies()
	if err != nil {
		return err
	}

	for {
		row, err := a.reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			facades.Log().Errorf("failed to read row %v", err)
			return err
		}

		rowCounter = rowCounter + 1

		var animal models.ImportAnimals
		animal.StateID = is.ID

		animal.LandingAt, err = time.Parse("02/01/2006", row[1])
		if err != nil {
			facades.Log().Errorf("failed to parse landing date: %v", err)
			animal.Reason = "A data de chegada do animal deve ser fornecida"
		}

		animal.Origin = row[2]
		_, err = helpers.GetGender(row[3])
		if err != nil {
			facades.Log().Errorf("failed to get correctly gender: %v", err)
			animal.Reason = "O Sexo do animal deve ser informado"
		}
		animal.Gender = row[3]

		_, err = helpers.GetAge(row[5])
		if err != nil {
			facades.Log().Errorf("failed to get correctly age: %v", err)
			animal.Reason = "A Idade do animal deve ser informada"
		}
		animal.Age = row[5]

		animal.BornDate, err = time.Parse("02/01/2006", row[4])
		if err != nil && animal.Age == "" {
			facades.Log().Errorf("failed to parse born date: %v", err)
			animal.Reason = "A data de nascimento do animal deve ser fornecida"
		}

		animal.Observation = row[6]
		animal.WasherCode = row[7]
		animal.MicrochipCode = row[8]

		if animal.LandingAt.Equal(time.Time{}) {
			animal.Reason = "A data de chegada do animal deve ser fornecida"
		}

		enclosure, ok := enclosures[normalize(row[13])]
		if ok {
			animal.EnclosureID = enclosure
		} else {
			animal.Reason = "O recinto está vazio, por isso não foi possível importar"
		}

		s, ok := species[normalize(row[9])]
		if ok {
			animal.SpeciesID = s
		} else {
			animal.Reason = "A espécie do animal está vazia, por isso não foi possível importar"
		}

		err = query.Create(&animal)
		if err != nil {
			continue
		}

		doneCounter = doneCounter + 1
	}

	facades.Log().Infof("Imported %d animals\n", doneCounter)
	return nil
}

func getEnclosuses() (map[string]string, error) {
	var enclosures []models.Enclosure
	query := facades.Orm().Query()

	err := query.Find(&enclosures)
	if err != nil {
		return map[string]string{}, err
	}

	mappedEnclosures := make(map[string]string)
	for _, e := range enclosures {
		mappedEnclosures[normalize(e.Identification)] = e.ID
	}

	return mappedEnclosures, nil
}

func getSpecies() (map[string]string, error) {
	var species []models.Species
	query := facades.Orm().Query()

	err := query.Find(&species)
	if err != nil {
		return map[string]string{}, err
	}

	mappedSpecies := make(map[string]string)
	for _, s := range species {
		mappedSpecies[normalize(s.CommonName)] = s.ID
	}

	return mappedSpecies, nil
}
