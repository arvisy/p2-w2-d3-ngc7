package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Store    Store  `json:"store" gorm:"foreignKey:StoreID"`
	StoreID  uint   `json:"-"`
}
