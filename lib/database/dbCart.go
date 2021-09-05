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

func CartCheck(customerId uint) uint {
	cart := model.Cart{
		CustomerID: customerId,
	}

	query := config.DB.Where("customer_id = ?", customerId).Find(&cart)

	if query.RowsAffected == 0 {
		config.DB.Create(&cart)
		CartCheck(customerId)
	}

	return cart.ID
}
