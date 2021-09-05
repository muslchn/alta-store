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

func GetProductById(id uint) (uint, uint, error) {
	product := model.Product{}
	query := config.DB.Where("id = ?", id).Find(&product)

	if err := query.Error; err != nil {
		return 0, 0, err
	}

	if query.RowsAffected == 0 {
		return 0, 0, nil
	}

	return product.Stock, product.Price, nil
}

func UpdateProductStockById(id, stock, qty uint) {
	product := model.Product{
		Stock: stock - qty,
	}

	config.DB.Where("product_id = ?", id).Updates(&product)
}
