package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	CustomerID uint
	Checkout   bool
	Customer   Customer
}

type CartItem struct {
	gorm.Model
	CartID    uint
	ProductID uint
	Price     uint
	Qty       uint
	Total     uint
	Cart      Cart
	Product   Product
}
