package models

import "time"

type MarkTypeStatus string

const (
	Washer    MarkTypeStatus = "wahser"
	Microchip MarkTypeStatus = "microchip"
)

type Animal struct {
	Model
	Name        string         `gorm:"column:name;default:null;" json:"name,omitempty"`
	MarkType    MarkTypeStatus `gorm:"column:mark_type;default:wahser;" json:"mark_type"`
	MarkNumber  string         `gorm:"column:mark_number" json:"mark_number"`
	LandingAt   time.Time      `gorm:"column:landing_at" json:"landing_at"`
	Age         int64          `gorm:"column:age" json:"age"`
	SpeciesId   string         `gorm:"column:species_id;type:uuid;not null" json:"species_id"`
	EnclosureId string         `gorm:"column:enclosure_id;type:uuid;not null" json:"enclosure_id"`
}

func (Animal) TableName() string {
	return "animals"
}
