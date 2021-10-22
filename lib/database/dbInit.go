package database

import (
	"alta-store/config"
	"alta-store/models"
)

func InitData() {
	var (
		category []models.Category
		product  []models.Product
	)

	if query := config.DB.Find(&category); query.RowsAffected == 0 {
		category = []models.Category{
			{Name: "Buku"},
			{Name: "Handphone"},
			{Name: "Laptop"},
		}

		config.DB.Create(&category)
	}

	if query := config.DB.Find(&product); query.RowsAffected == 0 {
		product = []models.Product{
			{
				CategoryID:  1,
				Name:        "Sapiens",
				Description: "Riwayat Singkat Umat Manusia ...",
				Stock:       3,
				Price:       115000,
				Weight:      700,
			},

			{
				CategoryID:  1,
				Name:        "Disruption",
				Description: "Disruptive Innovation ...",
				Stock:       3,
				Price:       116900,
				Weight:      500,
			},

			{
				CategoryID:  1,
				Name:        "Atomic Habits",
				Description: "Perubahan Kecil Yang Memberikan Hasil Luar Biasa ...",
				Stock:       3,
				Price:       86500,
				Weight:      500,
			},

			{
				CategoryID:  2,
				Name:        "iPhone 12 128GB",
				Description: "Garansi Resmi iBox Apple Indonesia ...",
				Stock:       3,
				Price:       13950000,
				Weight:      500,
			},

			{
				CategoryID:  2,
				Name:        "Samsung Galaxy A12 128GB",
				Description: "PLATFORM OS Android 10 ...",
				Stock:       3,
				Price:       2254000,
				Weight:      500,
			},

			{
				CategoryID:  2,
				Name:        "Oppo Reno6 128GB",
				Description: "Garansi Resmi Oppo 1 Tahun ...",
				Stock:       3,
				Price:       4405000,
				Weight:      500,
			},

			{
				CategoryID:  3,
				Name:        "MacBook Air M1",
				Description: "Macbook Terbaik ...",
				Stock:       3,
				Price:       15990000,
				Weight:      5000,
			},

			{
				CategoryID:  3,
				Name:        "ASUS ROG Zephyrus G15",
				Description: "Laptop Gaming Terbaik ...",
				Stock:       3,
				Price:       32999000,
				Weight:      6000,
			},

			{
				CategoryID:  3,
				Name:        "Acer Swift 3 Infinity 4",
				Description: "Laptop untuk Mahasiswa ...",
				Stock:       3,
				Price:       12499000,
				Weight:      5000,
			},
		}

		config.DB.Create(&product)
	}
}
