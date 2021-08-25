package models

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	OrderID        int       `json:"orderId" form:"orderId"`
	CustomerID     int       `json:"customerId" form:"customerId"`
	PaymentID      int       `json:"paymentId" form:"paymentId"`
	OrderTime      time.Time `json:"orderTime" form:"orderTime"`
	Timestamp      time.Time `json:"timestamp" form:"timestamp"`
	TransactStatus string    `json:"transactStatus" form:"transactStatus"`
	Fulfilled      bool      `json:"fulfilled" form:"fulfilled"`
	Deleted        bool      `json:"deleted" form:"deleted"`
	Paid           bool      `json:"paid" form:"paid"`
	PaymentTime    time.Time `json:"paymentTime" form:"paymentTime"`
}
