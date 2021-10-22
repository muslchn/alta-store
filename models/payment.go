package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CustomerID uint
	CheckoutID uint
	Paid       uint
	Customer   Customer
	Checkout   Checkout
}
