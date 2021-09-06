package database

import (
	"alta-store/config"
	"alta-store/model"
)

func GetTotal(cartId uint) (uint, uint, error) {
	var (
		cartItem   []model.CartItem
		totalItem  uint
		totalPrice uint
	)

	if err := config.DB.Where("cart_id = ?", cartId).Find(&cartItem).Error; err != nil {
		return 0, 0, err
	}

	for _, item := range cartItem {
		totalItem += item.Qty
		totalPrice += item.Total
	}

	return totalItem, totalPrice, nil
}

func CreatCheckout(CartId, totalItem, totalPrice uint) (interface{}, error) {
	checkout := model.Checkout{
		CartID:     CartId,
		TotalItem:  totalItem,
		TotalPrice: totalPrice,
		Paid:       false,
	}

	if err := config.DB.Create(&checkout).Error; err != nil {
		return nil, err
	}

	return checkout, nil
}
