package model

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name     string    `json:"name"`
	Products []Product `json:"products" gorm:"foreignKey:StoreID"`
}
