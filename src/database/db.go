package database

import (
	"ambassador/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Global variable has pointer to GORM DB
var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open(mysql.Open("docker:password@tcp(ambassadordb)/ambassador"), &gorm.Config{})

	if err != nil {
		panic("Gagal Terkoneksi")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.User{})
}
