package models

import (
	"time"

	"gorm.io/gorm"
)

type Shipments struct {
	gorm.Model
	ShipID   int       `json:"shipId" form:"shipId"`
	OrderID  int       `json:"orderId" form:"orderId"`
	ShipTime time.Time `json:"shipTime" form:"shipTime"`
}
