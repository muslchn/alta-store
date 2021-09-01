package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID uint `gorm:"primaryKey" json:"id"`
	// CategoryID  uint `json:"categoryId"`
	Name        string
	Description string
	Stock       uint
	Price       uint
	Weight      float32
	// Category    Category
}
