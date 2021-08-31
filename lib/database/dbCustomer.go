package database

import (
	"alta-store/config"
	"alta-store/models"
)

func Register(getCustomer models.Customer) (interface{}, error) {
	var customer models.Customer

	customer = getCustomer

	if err := config.DB.Create(&customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func GetDetailCustomers(customerId int) (interface{}, error) {
	var customer models.Customer

	if err := config.DB.Find(&customer, customerId).Error; err != nil {
		return nil, err
	}
	return customer, nil
}
