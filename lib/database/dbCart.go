package database

import (
	"alta-store/config"
	"alta-store/model"
)

func GetCart() (interface{}, error) {
	var products []model.Cart

	if err := config.DB.Preload("Product").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func AddToCart(id uint) (interface{}, error) {
	product := model.Cart{}

	if err := config.DB.Create(&model.CartItem{Product: model.Product{ID: id}}).Error; err != nil {
		return nil, err
	}

	return product, nil
}
