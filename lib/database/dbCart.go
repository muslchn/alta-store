package database

import (
	"alta-store/config"
	"alta-store/model"
	"strconv"
)

func GetCart(customerId uint) (interface{}, error) {
	var cartItem []model.CartItem
	cartId, _ := CartCheck(customerId)

	if err := config.DB.Where("cart_id = ?", cartId).Find(&cartItem).Error; err != nil {
		return nil, err
	}

	return cartItem, nil
}

func CartCheck(customerId uint) (uint, error) {
	var cart model.Cart

	query := config.DB.Where("customer_id = ? AND checkout = false", customerId).Find(&cart)

	if query.Error != nil {
		return 0, query.Error
	}

	if query.RowsAffected == 0 {
		cart = model.Cart{
			CustomerID: customerId,
			Checkout:   false,
		}

		config.DB.Create(&cart)
		CartCheck(customerId)
	}

	return cart.ID, nil
}

func AddCartItem(item map[string]string, customerId uint) ([]bool, interface{}, error) {
	// type productResponse struct {
	// 	Product bool
	// 	Stock   bool
	// 	Qty     bool
	// }

	productRes := []bool{false, false, false}

	cartId, _ := CartCheck(customerId)
	productId, _ := strconv.Atoi(item["product_id"])
	productValidation, productStock, productPrice, err := GetProductById(uint(productId))
	qty, _ := strconv.Atoi(item["qty"])
	total := productPrice * uint(qty)

	if err != nil {
		return nil, nil, err
	}

	if !productValidation {
		return productRes, nil, nil
	}

	if productStock == 0 {
		productRes[0] = true

		return productRes, nil, nil
	}

	if productStock < uint(qty) {
		productRes[0] = true
		productRes[1] = true

		return productRes, nil, nil
	}

	cartItem := model.CartItem{
		CartID:    cartId,
		ProductID: uint(productId),
		Price:     productPrice,
		Qty:       uint(qty),
		Total:     total,
	}

	if err := config.DB.Create(&cartItem).Error; err != nil {
		return nil, nil, err
	}

	UpdateProductStockById(uint(productId), productStock, uint(qty))

	return nil, cartItem, nil
}

func ValidateCartItem(cartItemId, customerId uint) (bool, map[string]uint, error) {
	var cartItem model.CartItem

	cartId, _ := CartCheck(customerId)
	query := config.DB.Where("id = ? AND cart_id = ?", cartItemId, cartId).Find(&cartItem)
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
