package database

import (
	"alta-store/config"
	"alta-store/model"
)

func Payment(customerId, checkoutId uint) (interface{}, error) {
	var (
		checkout model.Checkout
		payment  model.Payment
		paid     uint
	)

	query := config.DB.Where("id = ? AND customer_id = ?", checkoutId, customerId).First(&checkout)

	if query.Error != nil {
		return nil, query.Error
	}

	if query.RowsAffected == 0 {
		return nil, nil
	}

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

func LastPayment(customerId uint) (interface{}, error) {
	var lastPayment model.Payment

	if err := config.DB.Where("customer_id = ?", customerId).Last(&lastPayment).Error; err != nil {
		return nil, err
	}

	return lastPayment, nil
}

func PaymentHistory(customerId uint) (interface{}, error) {
	var paymentHistory []model.Payment

	if err := config.DB.Where("customer_id = ?", customerId).Find(&paymentHistory).Error; err != nil {
		return nil, err
	}

	return paymentHistory, nil
}

func PaymentById(id uint) (interface{}, error) {
	var paymentById model.Payment

	if err := config.DB.Where("id = ?", id).Find(&paymentById).Error; err != nil {
		return nil, err
	}

	return paymentById, nil
}
