package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID   uint
	Name string
}
