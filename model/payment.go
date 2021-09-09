package model

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CheckoutID uint
	Paid       uint
	Checkout   Checkout
}
