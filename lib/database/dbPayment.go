package database

import (
	"alta-store/config"
	"alta-store/model"
)

func Payment(checkoutId uint) (interface{}, error) {
	var (
		checkout   model.Checkout
		payment    model.Payment
		customerId uint
		paid       uint
	)

	if err := config.DB.Where("id = ?", checkoutId).Find(&checkout).Error; err != nil {
		return nil, err
	}

	customerId = checkout.CustomerID
	paid = checkout.TotalPrice

	payment = model.Payment{
		CustomerID: customerId,
		CheckoutID: checkoutId,
		Paid:       paid,
	}

	if err := config.DB.Create(&payment).Error; err != nil {
		return nil, err
	}

	return payment, nil
}
