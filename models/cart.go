package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	// CustomerID uint
	ProductID uint
	Quantity  uint
	Cost      uint
	TimeAdded time.Time
	// Customer  Customer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
