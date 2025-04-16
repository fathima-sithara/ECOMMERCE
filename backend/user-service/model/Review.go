package model

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	UserName  string `json:"user_name"`
	ProductID uint   `json:"product_id"`
	Rating    int    `json:"rating" gorm:"not null"`
	Comment   string `json:"comment"`
}
