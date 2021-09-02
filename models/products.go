package models

import "gorm.io/gorm"

type Products struct {
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
