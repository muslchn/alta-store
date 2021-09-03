package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	CustomersID uint
	Customers   Customers
}

type CartItems struct {
	gorm.Model
	CartID     uint
	ProductsID uint
	Qty        uint
	Cost       uint
	Cart       Cart
	Products   Products
}
