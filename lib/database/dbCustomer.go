package database

import (
	"alta-store/config"
	"alta-store/models"
)

func GetCustomers() (interface{}, error) {
	var customers []models.Customer

	if err := config.DB.Find(&customers).Error; err != nil {
		return nil, err
	}

	return customers, nil
}
