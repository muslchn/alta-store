package model

import (
	"time"

	"gorm.io/gorm"
)

type Shipment struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	OrderID uint
	Time    time.Time
	Order   Order
}
