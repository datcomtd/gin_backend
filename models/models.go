package models

import (
	"time"
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
	COMP = 1
	TSI  = 2
)

type Token struct {
	UpdatedAt time.Time
	String    string
}

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
