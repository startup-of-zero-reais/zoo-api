package models

import (
	"time"
)

type User struct {
	Model

	Name            string    `gorm:"column:name;not null;" json:"name,omitempty"`
	Avatar          string    `gorm:"column:avatar;default:null" json:"avatar,omitempty"`
	Email           string    `gorm:"column:email;not null" json:"email,omitempty"`
	EmailVerifiedAt time.Time `gorm:"column:email_verified_at" json:"email_verified_at,omitempty"`
}

func (User) TableName() string {
	return "users"
}

type Session struct {
	User        User   `json:"user"`
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}
