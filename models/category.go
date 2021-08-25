package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryID   int    `json:"categoryId" form:"categoryId"`
	CategoryName string `json:"categoryName" form:"categoryName"`
	Description  string `json:"descripttion" form:"description"`
	Active       bool   `json:"active" form:"active"`
}
