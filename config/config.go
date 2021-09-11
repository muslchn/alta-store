package config

import (
	"alta-store/model"
	"alta-store/secret"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": secret.DB_PASSWORD,
		"DB_Port":     "3306",
		"DB_Host":     "localhost",
		"DB_Name":     "alta-store",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"],
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&model.Cart{})
	DB.AutoMigrate(&model.Category{})
	DB.AutoMigrate(&model.Customer{})
	DB.AutoMigrate(&model.Payment{})
	DB.AutoMigrate(&model.Product{})
}
