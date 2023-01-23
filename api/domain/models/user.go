package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"unique;not null;size:255;"`
	Email     string         `json:"email"`
	Age       int            `json:"age"`
	Role      string         `json:"role"`
	Birthday  time.Time      `json:"birthday"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
