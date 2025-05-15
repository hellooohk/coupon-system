package repository

import (
	"coupon-system/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("coupons.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	err = DB.AutoMigrate(
		&models.Category{},
		&models.Medicine{},
		&models.Coupon{},
		&models.CouponMedicine{},
		&models.CouponCategory{},
	)
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}
}
