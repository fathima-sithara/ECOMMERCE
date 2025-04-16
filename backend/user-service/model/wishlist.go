package model

import (
	"gorm.io/gorm"
)

type Wishlist struct {
	gorm.Model
	UserID uint           `json:"user_id"` //foreignKey  user
	Items  []WishlistItem `gorm:"foreignKey:WishlistID"`
}

type WishlistItem struct {
	ID         uint `gorm:"primaryKey"`
	WishlistID uint `json:"wishlist_id"`
	ProductID  uint `json:"product_id"`
}
