package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	FirstName string `json:"firstName" form:"firstName"`
	LastName  string `json:"lastName" form:"lastName"`
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	// Token  string
	Address    string `json:"address" form:"address"`
	City       string `json:"city" form:"city"`
	State      string `json:"state" form:"state"`
	PostalCode uint   `json:"postalCode" form:"postalCode"`
	Phone      int    `json:"phone" form:"phone"`
	Email      string `json:"email" form:"email"`
}
