package database

import (
	"alta-store/config"
	"alta-store/models"
)

func GetCategory() (interface{}, error) {
	var category []models.Category

	if err := config.DB.Find(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func GetCategoryById(id uint) (interface{}, error) {
	var category models.Category

	if err := config.DB.Where("id = ?", id).Find(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}
