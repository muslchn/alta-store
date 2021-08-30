package database

import (
	"alta-store/config"
	"alta-store/models"
)

func GetCart() (interface{}, error) {
	var products []models.Cart

	if err := config.DB.Preload("Product").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func AddToCart(id uint) (interface{}, error) {
	product := models.Cart{}

	if err := config.DB.Create(&models.Cart{Product: models.Product{ID: id}}).Error; err != nil {
		return nil, err
	}

	return product, nil
}
