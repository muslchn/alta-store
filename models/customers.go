package models

import (
	"time"

	"gorm.io/gorm"
)

type Customers struct {
	gorm.Model
	CustomerID  int       `json:"id" form:"id"`
	FirstName   string    `json:"firstName" form:"firstName"`
	LastName    string    `json:"lastName" form:"lastName"`
	Username    string    `json:"username" form:"username"`
	Password    string    `json:"password" form:"password"`
	Token       string    `json:"token" form:"token"`
	Address     string    `json:"address" form:"address"`
	City        string    `json:"city" form:"city"`
	State       string    `json:"state" form:"state"`
	PostalCode  int       `json:"postalCode" form:"postalCode"`
	Phone       int       `json:"phone" form:"phone"`
	Email       string    `json:"email" form:"email"`
	TimeEntered time.Time `json:"timeEntered" form:"timeEntered"`
}
