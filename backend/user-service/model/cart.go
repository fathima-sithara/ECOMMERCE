package model

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID uint       `json:"user_id"`
	Items  []CartItem `gorm:"foreignKey:CartID"`
}

type CartItem struct {
	ID        uint    `gorm:"primaryKey"`
	CartID    uint    `json:"cart_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity" gorm:"default:1"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
}
