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

func GetProductsByCategory(categoryId uint) (interface{}, error) {
	var products []model.Product

	if err := config.DB.Where("category_id  = ?", categoryId).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductById(id uint) (bool, uint, uint, error) {
	var product model.Product

	query := config.DB.Where("id = ?", id).Find(&product)

	if err := query.Error; err != nil {
		return false, 0, 0, err
	}

	if query.RowsAffected == 0 {
		return false, 0, 0, nil
	}

	return true, product.Stock, product.Price, nil
}

func UpdateProductStockById(id, stock, qty uint) {
	product := model.Product{
		Stock: stock - qty,
	}

	config.DB.Where("product_id = ?", id).Updates(&product)
}

func ReturnStock(productId, qty uint) {
	product := model.Product{}
	stock := 0
	query := config.DB.Where("id = ?", productId).Find(&product)

	if query.RowsAffected > 0 {
		stock = int(product.Stock)
		returnStock := model.Product{
			Stock: uint(stock) + qty,
		}

		config.DB.Where("id = ?", productId).Updates(&returnStock)
	}
}
