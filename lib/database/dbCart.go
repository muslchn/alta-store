package database

import (
	"alta-store/config"
	"alta-store/model"
	"strconv"
)

func GetCart(customerId uint) (interface{}, error) {
	var cartItem []model.CartItem
	cartId := CartCheck(customerId)

	if err := config.DB.Where("cart_id = ?", cartId).Find(&cartItem).Error; err != nil {
		return nil, err
	}

	return cartItem, nil
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

func AddCartItem(payloadData map[string]string, customerId uint) (interface{}, error) {
	cartId := CartCheck(customerId)
	productId, _ := strconv.Atoi(payloadData["product_id"])
	productStock, productPrice, err := GetProductById(uint(productId))
	qty, _ := strconv.Atoi(payloadData["qty"])
	total := productPrice * uint(qty)

	if err != nil {
		return nil, err
	}

	if productStock < uint(qty) {
		return false, nil
	}

	cartItem := model.CartItem{
		CartID:    cartId,
		ProductID: uint(productId),
		Price:     productPrice,
		Qty:       uint(qty),
		Total:     total,
	}

	if err := config.DB.Create(&cartItem).Error; err != nil {
		return nil, err
	}

	UpdateProductStockById(uint(productId), productStock, uint(qty))

	return cartItem, nil
}
