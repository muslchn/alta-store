package database

import (
	"alta-store/config"
	"alta-store/model"
)

func GetCategory() (interface{}, error) {
	var category []model.Category

	if err := config.DB.Find(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func GetCategoryById(id int) (interface{}, error) {
	var category []model.Category

	if err := config.DB.Where("id = ?", id).Find(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}
