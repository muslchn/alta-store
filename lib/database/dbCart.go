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
		Checkout:   false,
	}

	query := config.DB.Where("customer_id = ? AND checkout = false", customerId).Find(&cart)

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

func ValidateCartItem(cartItemId, customerId uint) (bool, map[string]uint, error) {
	cartId := CartCheck(customerId)
	cartItem := model.CartItem{}
	query := config.DB.Where("id = ?", cartItemId).Find(&cartItem)
	returnItem := make(map[string]uint)

	if query.Error != nil {
		return false, returnItem, query.Error
	} else if cartItem.ID != cartId {
		return false, returnItem, nil
	}

	returnItem["product_id"] = cartItem.ProductID
	returnItem["qty"] = cartItem.Qty

	return true, returnItem, nil
}

func DeleteCartItem(id uint) (interface{}, error) {
	var cartItem []model.CartItem

	if err := config.DB.Where("id = ?", id).Delete(&cartItem).Error; err != nil {
		return nil, err
	}

	return true, nil
}

func ChangeCartStatus(id uint) error {
	cart := model.Cart{
		Checkout: true,
	}

	if err := config.DB.Where("id = ?", id).Updates(&cart).Error; err != nil {
		return err
	}

	return nil
}
