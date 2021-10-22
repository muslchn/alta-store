package database

import (
	"alta-store/config"
	"alta-store/models"
)

func GetProducts() (interface{}, error) {
	var products []models.Product

	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductsByCategory(categoryId uint) (interface{}, error) {
	var products []models.Product

	query := config.DB.Where("category_id  = ?", categoryId).Find(&products)

	if query.Error != nil {
		return nil, query.Error
	}

	if query.RowsAffected == 0 {
		return nil, nil
	}

	return products, nil
}

func GetProductById(id uint) (bool, uint, uint, error) {
	var product models.Product

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
	product := models.Product{
		Stock: stock - qty,
	}

	config.DB.Where("product_id = ?", id).Updates(&product)
}

func ReturnStock(productId, qty uint) error {
	var (
		product models.Product
		stock   uint
	)

	query := config.DB.Where("id = ?", productId).First(&product)

	if query.Error != nil {
		return query.Error
	}

	if query.RowsAffected != 0 {
		stock = product.Stock
		returnStock := models.Product{
			Stock: stock + qty,
		}

		config.DB.Where("id = ?", productId).Updates(&returnStock)
	}

	return nil
}
