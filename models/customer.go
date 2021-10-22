package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	FirstName  string `gorm:"size:255;not null" json:"firstName" form:"firstName"`
	LastName   string `gorm:"size:255" json:"lastName" form:"lastName"`
	Username   string `gorm:"size:255;not null;unique" json:"username" form:"username"`
	Email      string `gorm:"size:100;not null;unique" valid:"email" json:"email" form:"email"`
	Password   string `gorm:"size:100;not null" json:"password" form:"password"`
	Phone      string `gorm:"size:15;not null" json:"phone" form:"phone"`
	Address    string `gorm:"size:255;not null" json:"address" form:"address"`
	City       string `gorm:"size:100;not null" json:"city" form:"city"`
	State      string `gorm:"size:100;not null" json:"state" form:"state"`
	PostalCode string `gorm:"5;not null" json:"postalCode" form:"postalCode"`
	Token      string `json:"token" form:"token"`
}
