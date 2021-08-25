package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	CartID     int       `json:"cartId" form:"cartId"`
	CustomerID int       `json:"customerId" form:"customerId"`
	ProductID  int       `json:"productId" form:"productId"`
	Quantity   int       `json:"quantity" form:"quantity"`
	Cost       int       `json:"cost" form:"cost"`
	TimeAdded  time.Time `json:"timeAdded" form:"timeAdded"`
}
