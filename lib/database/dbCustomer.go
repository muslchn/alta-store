package database

import (
	"alta-store/config"
	"alta-store/models"
)

func Register() (interface{}, error) {
	var customer models.Customer

	if err := config.DB.Find(&customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}
