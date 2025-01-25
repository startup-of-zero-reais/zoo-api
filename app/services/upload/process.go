package upload

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
)

func (e uploadImpl) Process(filePath string) error {

	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)

	if err != nil {
		facades.Log().Errorf("failed to upload process %v", err)
		return err
	}

	csvReader := csv.NewReader(file)
	csvReader.Read()

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			facades.Log().Errorf("failed to read row %v", err)
			return err
		}

		name := row[0]
		common_name := row[1]
		scientific_name := row[2]
		landing_at := row[3]
		origin := row[4]
		gender := row[5]
		born_date := row[6]
		age := row[7]
		observation := row[8]
		washer_code := row[9]
		microchip_code := row[10]
		taxonomic_order := row[11]
		kind := row[12]
		identification := row[13]

		var ce requests.CreateEnclosure
		ce.Identification = identification

		enclosure, err := e.EnclosureService.Create(ce)
		if err != nil {
			facades.Log().Errorf("failed to create enclosure on upload %v", err)
			return err
		}

		var cs requests.CreateSpecies
		cs.CommonName = common_name
		cs.Kind = kind
		cs.ScientificName = scientific_name
		cs.TaxonomicOrder = taxonomic_order

		species, err := e.SpeciesService.Create(cs)
		if err != nil {
			facades.Log().Errorf("failed to create species on upload %v", err)
			return err
		}

		var ca requests.CreateAnimal
		ca.Name = name
		ca.Age = age
		ca.BornDate = born_date
		ca.Gender = gender
		ca.MicrochipCode = microchip_code
		ca.WasherCode = washer_code
		ca.Observation = observation
		ca.LandingAt = landing_at
		ca.Origin = origin
		ca.EnclosureID = enclosure.ID
		ca.SpeciesID = species.ID

		_, err = e.AnimalService.Create(ca)
		if err != nil {
			facades.Log().Errorf("failed to create animal on upload %v", err)
			return err
		}

		fmt.Println("animal created by upload", name)
		time.Sleep(time.Second * 3)
	}

	os.Remove(filePath)

	return nil
}
