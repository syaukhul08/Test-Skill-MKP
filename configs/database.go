package configs

import (
	"GO-TEST-MKP/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConfigDB() {

	dsn := "host=localhost password=jepara08 user=postgres dbname=ticketing port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Terminal{})
	DB = db
}