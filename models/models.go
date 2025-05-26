package models

import (
	"time"

	"github.com/lib/pq"
)

const (
	President uint = 1
	Vice_tsi       = 2
	Vice_comp 	   = 3
	Secretary      = 4
	Treasurer      = 5
	Director       = 6
	Other 		   = 7
)

const (
	COMP  uint = 1 
	TSI        = 2
	ELET       = 3
	CIVIL      = 4
	EBB        = 5
	TPQ		   = 6
	MAT 	   = 7
	OTHER	   = 8
)

type User struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time

	ID uint `json:"id" gorm:"primaryKey"`

	Token_UpdatedAt time.Time
	Token           string
	Password        string

	Username string `json:"name"`
	Email    string `json:"email"`

	Role   uint   `json:"role"`
	Course uint   `json:"course"`
	RA     string `json:"ra"`
}

type Document struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time

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
	CreatedAt  time.Time
	UpdatedAt  time.Time

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

type Booking struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time

	ID 	uint `json:"id" gorm:"primaryKey"`

	TimestampStart time.Time `json:"time-start"`
	TimestampEnd   time.Time `json:"time-end"`
	Description    string    `json:"description"`

	Day 		string
	Title 		string 	`json:"title"`
	Location 	string 	`json:"location"`
	Private 	bool 	`json:"private"`

	Username string `json:"username"`
	Role     uint   `json:"role"`
	Course   uint   `json:"course"`

	CreatedBy     string `json:"created-by"`
	LastUpdatedBy string `json:"last-updated-by"`
}
