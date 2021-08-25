package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	ProductID          int     `json:"productId" form:"productId"`
	CategoryID         int     `json:"categoryId" form:"categoryId"`
	SKU                string  `json:"sku" form:"sku"`
	ProductName        string  `json:"productName" form:"productName"`
	ProductDescription string  `json:"productDescripttion" form:"productDescriptiton"`
	Stock              int     `json:"stock" form:"stock"`
	Price              int     `json:"price" form:"price"`
	Size               int     `json:"size" form:"size"`
	Color              string  `json:"color" form:"color"`
	Weight             float32 `json:"weight" form:"weight"`
}
