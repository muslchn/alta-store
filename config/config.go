package config

import (
	"alta-store/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "123ABC4d.",
		"DB_Port":     "3306",
		"DB_Host":     "localhost",
		"DB_Name":     "alta",
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

	/*
		DB.Create(&models.Category{
			ID:   1,
			Name: "Buku",
		})

		DB.Create(&models.Category{
			ID:   2,
			Name: "Handphone",
		})

		DB.Create(&models.Category{
			ID:   3,
			Name: "Laptop",
		})

		DB.Create(&models.Product{
			ID:          11,
			CategoryID:  1,
			Name:        "Sapiens",
			Description: "Riwayat Singkat Umat Manusia ...",
			Stock:       3,
			Price:       115000,
			Weight:      700,
		})

		DB.Create(&models.Product{
			ID:          12,
			CategoryID:  1,
			Name:        "Disruption",
			Description: "Disruptive Innovation ...",
			Stock:       3,
			Price:       116900,
			Weight:      500,
		})

		DB.Create(&models.Product{
			ID:          13,
			CategoryID:  1,
			Name:        "Atomic Habits",
			Description: "Perubahan Kecil Yang Memberikan Hasil Luar Biasa ...",
			Stock:       3,
			Price:       86500,
			Weight:      500,
		})

		DB.Create(&models.Product{
			ID:          21,
			CategoryID:  2,
			Name:        "iPhone 12 128GB",
			Description: "Garansi Resmi iBox Apple Indonesia ...",
			Stock:       3,
			Price:       13950000,
			Weight:      500,
		})

		DB.Create(&models.Product{
			ID:          22,
			CategoryID:  2,
			Name:        "Samsung Galaxy A12 128GB",
			Description: "PLATFORM OS Android 10 ...",
			Stock:       3,
			Price:       2254000,
			Weight:      500,
		})

		DB.Create(&models.Product{
			ID:          23,
			CategoryID:  2,
			Name:        "Oppo Reno6 128GB",
			Description: "Garansi Resmi Oppo 1 Tahun ...",
			Stock:       3,
			Price:       4405000,
			Weight:      500,
		})

		DB.Create(&models.Product{
			ID:          31,
			CategoryID:  3,
			Name:        "MacBook Air M1",
			Description: "Macbook Terbaik ...",
			Stock:       3,
			Price:       15990000,
			Weight:      5000,
		})

		DB.Create(&models.Product{
			ID:          32,
			CategoryID:  3,
			Name:        "ASUS ROG Zephyrus G15",
			Description: "Laptop Gaming Terbaik ...",
			Stock:       3,
			Price:       32999000,
			Weight:      6000,
		})

		DB.Create(&models.Product{
			ID:          33,
			CategoryID:  3,
			Name:        "Acer Swift 3 Infinity 4",
			Description: "Laptop untuk Mahasiswa ...",
			Stock:       3,
			Price:       12499000,
			Weight:      5000,
		})

		DB.Create(&models.Payment{
			ID:   1,
			Type: "LinkAja",
		})

		DB.Create(&models.Payment{
			ID:   2,
			Type: "OVO",
		})

		DB.Create(&models.Payment{
			ID:   3,
			Type: "GoPay",
		})
	*/
}

func InitMigrate() {
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.Customer{})
	DB.AutoMigrate(&models.OrderDetail{})
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.Payment{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Shipment{})
}
