package models

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	ID             uint `gorm:"primaryKey"`
	CustomerID     uint
	PaymentID      uint
	OrderTime      time.Time
	Timestamp      time.Time
	TransactStatus string
	Fulfilled      bool
	Deleted        bool
	Customers      Customers
	Payment        Payment
}
