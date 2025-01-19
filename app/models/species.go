package models

type Species struct {
	Model

	CommonName    string `gorm:"column:common_name;default:null" json:"common_name"`
	CientificName string `gorm:"column:cientific_name;default:null" json:"cientific_name"`
}

func (Species) TableName() string {
	return "species"
}
