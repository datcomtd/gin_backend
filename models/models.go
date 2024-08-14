package models

import (
	"time"

	"github.com/lib/pq"
)

const (
	President uint = 1
	Vice_CC        = 2
	Vice_TSI       = 3
	Secretary      = 4
	Treasurer      = 5
	Director       = 6
)

const (
	COMP uint = 1
	TSI       = 2
)

type User struct {
	CreatedAt time.Time
	UpdateAt  time.Time

	Token_UpdatedAt time.Time
	Token           string
	Password        string

	Username string `json:"name" gorm:"primaryKey;unique"`
	Email    string `json:"email"`

	Role   uint `json:"role"`
	Course uint `json:"course"`
}

type Document struct {
	CreatedAt time.Time
	UpdateAt  time.Time

	ID  uint `json:"id" gorm:"primaryKey"`
	Key string

	Filename    string
	Title       string `json:"title"`
	Description string `json:"description"`

	Source   string `json:"source"`
	Category string `json:"category"`

	CreatedBy     string `json:"created-by"`
	LastUpdatedBy string `json:"last-updated-by"`
}

type Product struct {
	CreatedAt time.Time
	UpdateAt  time.Time

	ID uint `json:"id" gorm:"primaryKey"`

	Count  uint           `json:"count"`
	Photos pq.StringArray `json:"photos" gorm:"type:text[]"`

	Title       string `json:"title"`
	Description string `json:"description"`

	InStock bool `json:"in-stock"`

	CreatedBy     string `json:"created-by"`
	LastUpdatedBy string `json:"last-updated-by"`
}
