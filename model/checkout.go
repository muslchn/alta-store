package model

import "gorm.io/gorm"

type Checkout struct {
	gorm.Model
	CustomerID uint
	CartID     uint `json:"cartId" form:"cartId"`
	TotalItem  uint
	TotalPrice uint
	Paid       bool
	Customer   Customer
	Cart       Cart
}
