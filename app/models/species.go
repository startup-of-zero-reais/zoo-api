package models

type Species struct {
	Model

	CommonName     string `gorm:"column:common_name;default:null" json:"common_name"`
	CientificName  string `gorm:"column:cientific_name;default:null" json:"cientific_name"`
	Kind           string `gorm:"column:kind;default:null" json:"kind"`
	TaxonomicOrder string `gorm:"column:taxonomic_order;default:null" json:"taxonomic_order"`
}

func (Species) TableName() string {
	return "species"
}
