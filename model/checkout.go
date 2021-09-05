package model

import "gorm.io/gorm"

type Checkout struct {
	gorm.Model
	CartID     uint
	TotalItem  uint
	TotalPrice uint
	Cart       Cart
}
