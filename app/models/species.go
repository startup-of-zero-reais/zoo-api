package models

type Species struct {
	Model

	CommonName     string `gorm:"column:common_name;default:null" json:"common_name"`
	ScientificName string `gorm:"column:scientific_name;default:null" json:"scientific_name"`
	Kind           string `gorm:"column:kind;default:null" json:"kind"`
	TaxonomicOrder string `gorm:"column:taxonomic_order;default:null" json:"taxonomic_order"`
}

func (Species) TableName() string {
	return "species"
}
