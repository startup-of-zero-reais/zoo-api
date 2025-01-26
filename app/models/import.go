package models

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
}

func (ImportSpecies) TableName() string {
	return "import_species"
}

type ImportAnimals struct {
	// TODO : add fields
}

func (ImportAnimals) TableName() string {
	return "import_animals"
}
