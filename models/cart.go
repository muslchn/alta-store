package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	CustomerID uint
	Customer   Customer
}

type CartItem struct {
	gorm.Model
	CartID    uint
	ProductID uint
	Qty       uint
	Cost      uint
	Cart      Cart
	Product   Product
}
