package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderDetails struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	ProductID uint
	Price     uint
	Quantity  uint
	Total     uint
	Size      uint
	Color     string
	Fulfilled bool
	BillTime  time.Time
	Order     Orders
	Products  Products
}
