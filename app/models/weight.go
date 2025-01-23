package models

type Weight struct {
	Model
	Weight   float64 `gorm:"column:weight;not null;" json:"weight"`
	AnimalID string  `gorm:"column:animal_id;not null;" json:"animal_id"`
	UserID   string  `gorm:"column:user_id;not null;" json:"user_id"`
	Animal   *Animal `json:"animal,omitempty"`
	User     *User   `json:"user,omitempty"`
}

func (Weight) TableName() string {
	return "weight_history"
}
