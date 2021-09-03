package database

import (
	"alta-store/config"
	"alta-store/model"
)

func GetProducts() (interface{}, error) {
	var products []model.Product

	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductsByCategory(id int) (interface{}, error) {
	var products []model.Product

	if err := config.DB.Where("category_id  = ?", id).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

/* func AddProduct() (interface{}, error) {
	var products []models.Product

	if err := config.DB.Select(); err != nil {
		return nil, err
	}
} */
