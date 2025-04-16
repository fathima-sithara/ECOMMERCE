package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	UserID  uint   `json:"user_id" validate:"required"`
	OrderID uint   `json:"order_id" validate:"required"`
	Street  string `json:"street" gorm:"not null" validate:"required"`
	City    string `json:"city" gorm:"not null" validate:"required"`
	State   string `json:"state" gorm:"not null" validate:"required"`
	ZipCode string `json:"zip_code" gorm:"not null" validate:"required"`
	Country string `json:"country" gorm:"not null" validate:"required"`
}
type GetOrderdetils struct {
	gorm.Model
	UserID  uint   `json:"user_id" validate:"required"`
	OrderID uint   `json:"order_id" validate:"required"`
	Street  string `json:"street" gorm:"not null" validate:"required"`
	City    string `json:"city" gorm:"not null" validate:"required"`
	State   string `json:"state" gorm:"not null" validate:"required"`
	ZipCode string `json:"zip_code" gorm:"not null" validate:"required"`
	Country string `json:"country" gorm:"not null" validate:"required"`
	Method  string `json:"method" validate:"required"`
}
