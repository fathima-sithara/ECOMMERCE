package models

import (
	"gorm.io/gorm"
)

type Wishlist struct {
	gorm.Model
	UserID uint           `json:"user_id"`
	Items  []WishlistItem `gorm:"foreignKey:WishlistID"`
}

type WishlistItem struct {
	ID         uint `gorm:"primaryKey"`
	WishlistID uint `json:"wishlist_id"`
	ProductID  uint `json:"product_id"`
}
