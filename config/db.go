package config

import (
	"ngc7/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=ngc7 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	db.AutoMigrate(&model.Store{})
	db.AutoMigrate(&model.Product{}, &model.User{})

	return db
}
