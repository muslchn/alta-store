package database

import (
	"alta-store/config"
	"alta-store/model"
)

func GetTotal(cartId uint) (uint, uint, error) {
	var (
		cartItems  []model.CartItem
		totalItem  uint
		totalPrice uint
	)

	if err := config.DB.Where("cart_id = ?", cartId).Find(&cartItems).Error; err != nil {
		return 0, 0, err
	}

	for _, item := range cartItems {
		totalItem += item.Qty
		totalPrice += item.Total
	}

	return totalItem, totalPrice, nil
}

func CreateCheckout(customerId, cartId, totalItem, totalPrice uint) (interface{}, error) {
	// var cart model.Cart

	// if err := config.DB.Where("id = ?", cartId).Find(&cart).Error; err != nil {
	// 	return nil, err
	// }

	// customerId := cart.CustomerID

	checkout := model.Checkout{
		CustomerID: customerId,
		CartID:     cartId,
		TotalItem:  totalItem,
		TotalPrice: totalPrice,
		Paid:       false,
	}

	if err := config.DB.Create(&checkout).Error; err != nil {
		return nil, err
	}

	return checkout, nil
}

func GetCheckout(customerId uint) (interface{}, error) {
	var checkout []model.Checkout

	if err := config.DB.Where("customer_id = ? AND paid = false", customerId).Find(&checkout).Error; err != nil {
		return nil, err
	}

	return checkout, nil
}

func GetCheckoutById(id uint) (interface{}, error) {
	var checkout model.Checkout

	if err := config.DB.Where("id = ?", id).Find(&checkout).Error; err != nil {
		return nil, err
	}

	return checkout, nil
}
