package models

type Enclosure struct {
	Model
	Identification string `gorm:"column:identification;default:null;" json:"identification,omitempty"`
}

func (Enclosure) TableName() string {
	return "enclosures"
}
