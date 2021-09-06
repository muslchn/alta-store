package model

import "gorm.io/gorm"

type Checkout struct {
	gorm.Model
	CartID     uint `json:"cartId" form:"cartId"`
	TotalItem  uint
	TotalPrice uint
	Paid       bool
	Cart       Cart
}
