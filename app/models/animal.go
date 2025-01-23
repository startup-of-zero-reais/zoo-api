package models

import "time"

type Animal struct {
	Model
	Name          string    `gorm:"column:name;default:null;" json:"name,omitempty"`
	WasherCode    string    `gorm:"column:washer_code;default:null" json:"washer_code,omitempty"`
	MicrochipCode string    `gorm:"column:microchip_code;default:null" json:"microchip_code,omitempty"`
	LandingAt     time.Time `gorm:"column:landing_at" json:"landing_at"`
	Origin        string    `gorm:"column:origin" json:"origin"`
	BornDate      time.Time `gorm:"column:born_date" json:"born_date"`
	Age           string    `gorm:"column:age;default:null;" json:"age,omitempty"`
	Gender        string    `gorm:"column:gender;default:null;" json:"gender,omitempty"`
	Observation   string    `gorm:"column:observation;default:null;" json:"observation,omitempty"`
	SpeciesID     string    `gorm:"column:species_id;type:uuid;not null" json:"species_id"`
	EnclosureID   string    `gorm:"column:enclosure_id;type:uuid;not null" json:"enclosure_id"`

	Enclosure *Enclosure `json:"enclosure,omitempty"`
	Species   *Species   `json:"species,omitempty"`
}

func (Animal) TableName() string {
	return "animals"
}
