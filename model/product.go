package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required,min=50"`
	ImageURL    string `json:"image_url"`
	Price       int    `json:"price" binding:"required,min=1000"`
	StoreID     uint   `json:"store_id"`
	Store       Store  `json:"-" gorm:"foreignKey:StoreID"`
}
