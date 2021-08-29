package database

import (
	"alta-store/config"
	"alta-store/models"
)

func AddToCart(id uint) (interface{}, error) {
	product := models.Cart{ProductID: id}

	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}
