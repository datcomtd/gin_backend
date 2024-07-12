package models

import (
	"time"
)

type Document struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Title       string `json:"title"`
	File        string `json:"file"` // file path
	Description string `json:"description"`

	Category string `json:"category"`
	Active   bool   `json:"active"`
}
