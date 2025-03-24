package models

import (
	"time"

	"github.com/lib/pq"
)

const (
	President uint = 1
	Vice           = 2
	Secretary      = 3
	Treasurer      = 4
	Director       = 5
)

const (
	COMP  uint = 1
	TSI        = 2
	ELET       = 3
	CIVIL      = 4
	EBB        = 5
)

type User struct {
	CreatedAt time.Time
	UpdateAt  time.Time

	Token_UpdatedAt time.Time
	Token           string
	Password        string

	Username string `json:"name" gorm:"primaryKey"`
	Email    string `json:"email"`

	Role   uint   `json:"role"`
	Course uint   `json:"course"`
	RA     string `json:"ra"`
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

	Count  int            `json:"count"`
	Photos pq.StringArray `json:"photo" gorm:"type:text[]"`

	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`

	Price   float64 `json:"price"`
	InStock bool    `json:"in-stock"`

	CreatedBy     string `json:"created-by"`
	LastUpdatedBy string `json:"last-updated-by"`
}
