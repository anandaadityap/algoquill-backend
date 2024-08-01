package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name         string         `json:"name"`
	Biography    string         `json:"biography"`
	Email        string         `gorm:"unique" json:"email"`
	Password     string         `json:"password"`
	RefreshToken string         `json:"refresh_token"`
}
