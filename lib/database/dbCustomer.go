package database

import (
	"alta-store/config"
	"alta-store/middleware"
	"alta-store/model"
)

func Register(getCustomer *model.Customer) (interface{}, error) {

	customer := *getCustomer

	if err := config.DB.Create(&customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func GetDetailCustomers(customerId int) (interface{}, error) {
	var customer model.Customer

	if err := config.DB.Find(&customer, customerId).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func Login(email, password string) (interface{}, string, string, error) {
	var customer model.Customer

	if err := config.DB.Where("email = ? AND password = ?", email, password).First(&customer).Error; err != nil {
		return nil, "", "", err
	}

	token, err := middleware.CreateToken(int(customer.ID))

	if err != nil {
		return nil, "", "", err
	}

	customer.Token = token

	if err := config.DB.Save(&customer).Error; err != nil {
		return nil, "", "", err
	}

	return customer, customer.FirstName, customer.Token, nil
}
