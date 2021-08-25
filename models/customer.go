package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID          uint
	FirstName   string
	LastName    string
	Username    string
	Password    string
	Token       string
	Address     string
	City        string
	State       string
	PostalCode  uint
	Phone       int
	Email       string
	TimeEntered time.Time
}
