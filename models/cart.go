package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID         uint
	CustomerID uint
	ProductID  uint
	Quantity   uint
	Cost       uint
	TimeAdded  time.Time
	Customer   Customer
	Product    Product
}
