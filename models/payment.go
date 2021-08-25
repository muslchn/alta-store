package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	PaymentID   int    `json:"paymentId" form:"paymentId"`
	PaymentType string `json:"paymentType" form:"paymentType"`
	Allowed     bool   `json:"allowed" form:"allowed"`
}
