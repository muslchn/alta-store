package database

import (
	"alta-store/config"
	"alta-store/model"
)

func InitData() {
	var category []model.Category

	if query := config.DB.Find(&category); query.RowsAffected == 0 {
		category = []model.Category{
			{Name: "Buku"},
			{Name: "Handphone"},
			{Name: "Laptop"},
		}

		config.DB.Create(&category)
	}
}
