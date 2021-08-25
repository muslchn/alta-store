package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderDetails struct {
	gorm.Model
	OrderDetailID int       `json:"orderDetailId" form:"orderDetailId"`
	OrderID       int       `json:"orderId" form:"orderId"`
	ProductID     int       `json:"productId" form:"productId"`
	Price         int       `json:"price" form:"price"`
	Quantity      int       `json:"quantity" form:"quantity"`
	Total         int       `json:"total" form:"total"`
	Size          int       `json:"size" form:"size"`
	Color         string    `json:"color" form:"color"`
	Fulfilled     bool      `json:"fulfilled" form:"fulfilled"`
	BillTime      time.Time `json:"billTime" form:"billTime"`
}
