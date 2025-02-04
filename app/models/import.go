package models

import "time"

type ImportStatus struct {
	ID       string `gorm:"column:id;primaryKey;not null;default:uuid_generate_v4()" json:"id"`
	Filename string `gorm:"column:filename;default:null" json:"filename"`
	State    string `gorm:"column:state" json:"state"`
	UploadID string `gorm:"column:upload_id;default:null" json:"upload_id"`
}

func (ImportStatus) TableName() string {
	return "import_state"
}

type ImportEnclosure struct {
	ID             string `gorm:"column:id;primaryKey;not null;default:uuid_generate_v4()" json:"id"`
	Identification string `gorm:"column:identification;default:null" json:"identification"`
	Reason         string `gorm:"column:reason;default:null" json:"reason"`
	StateID        string `gorm:"column:state_id" json:"state_id"`
}

func (ImportEnclosure) TableName() string {
	return "import_enclosures"
}

type ImportSpecies struct {
	ID             string `gorm:"column:id;primaryKey;not null;default:uuid_generate_v4()" json:"id"`
	CommonName     string `gorm:"column:common_name;default:null;" json:"common_name"`
	ScientificName string `gorm:"column:scientific_name;default:null" json:"scientific_name"`
	Order          string `gorm:"column:taxonomic_order;default:null" json:"taxonomic_order"`
	Kind           string `gorm:"column:kind;default:null" json:"kind"`
	Reason         string `gorm:"column:reason;default:null" json:"reason"`
	StateID        string `gorm:"column:state_id" json:"state_id"`
}

func (ImportSpecies) TableName() string {
	return "import_species"
}

type ImportAnimals struct {
	ID            string    `gorm:"column:id;primaryKey;not null;default:uuid_generate_v4()" json:"id"`
	Name          string    `gorm:"column:name;default:null;" json:"name"`
	WasherCode    string    `gorm:"column:washer_code;default:null" json:"washer_code"`
	MicrochipCode string    `gorm:"column:microchip_code;default:null" json:"microchip_code"`
	LandingAt     time.Time `gorm:"column:landing_at;default:null" json:"landing_at"`
	Origin        string    `gorm:"column:origin;default:null" json:"origin"`
	BornDate      time.Time `gorm:"column:born_date;default:null" json:"born_date"`
	Age           string    `gorm:"column:age;default:null;" json:"age"`
	Gender        string    `gorm:"column:gender;default:null;" json:"gender"`
	Observation   string    `gorm:"column:observation;default:null;" json:"observation"`
	SpeciesID     *string   `gorm:"column:species_id;type:uuid;default:null" json:"species_id"`
	EnclosureID   *string   `gorm:"column:enclosure_id;type:uuid;default:null" json:"enclosure_id"`
	Reason        string    `gorm:"column:reason;default:null" json:"reason"`
	StateID       string    `gorm:"column:state_id" json:"state_id"`
}

func (ImportAnimals) TableName() string {
	return "import_animals"
}
