package database

import (
	"alta-store/config"
	"alta-store/middlewares"
	"alta-store/models"
)

func Register(getCustomer *models.Customers) (interface{}, error) {
	var customer models.Customers

	customer = *getCustomer

	if err := config.DB.Create(&customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func GetDetailCustomers(customerId int) (interface{}, error) {
	var customer models.Customers

	if err := config.DB.Find(&customer, customerId).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func LoginCustomers(customer *models.Customers) (interface{}, error) {

	var err error
	if err = config.DB.Where("email = ? AND password = ?", customer.Email, customer.Password).First(customer).Error; err != nil {
		return nil, err
	}

	customer.Token, err = middlewares.CreateToken(int(customer.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}
