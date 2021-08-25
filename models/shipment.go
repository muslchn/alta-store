package models

import (
	"time"

	"gorm.io/gorm"
)

type Shipment struct {
	gorm.Model
	ID      uint
	OrderID uint
	Time    time.Time
	Order   Order
}
