package models

import "time"

type MarkTypeStatus string

const (
	Washer    MarkTypeStatus = "washer"
	Microchip MarkTypeStatus = "microchip"
)

type Animal struct {
	Model
	Name        string         `gorm:"column:name;default:null;" json:"name,omitempty"`
	MarkType    MarkTypeStatus `gorm:"column:mark_type;default:wahser;" json:"mark_type"`
	MarkNumber  string         `gorm:"column:mark_number" json:"mark_number"`
	LandingAt   time.Time      `gorm:"column:landing_at" json:"landing_at"`
	Origin      string         `gorm:"column:origin" json:"origin"`
	Age         time.Time      `gorm:"column:age" json:"age"`
	SpeciesID   string         `gorm:"column:species_id;type:uuid;not null" json:"species_id"`
	EnclosureID string         `gorm:"column:enclosure_id;type:uuid;not null" json:"enclosure_id"`

	Enclosure *Enclosure `json:"enclosure,omitempty"`
	Species   *Species   `json:"species,omitempty"`
}

func (Animal) TableName() string {
	return "animals"
}
