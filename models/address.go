package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID  uint   `json:"user_id" validate:"required"`
	OrderID uint   `json:"order_id" validate:"required"`
	Street  string `gorm:"not null" json:"street" validate:"required"`
	City    string `gorm:"not null" json:"city" validate:"required"`
	State   string `gorm:"not null" json:"state" validate:"required"`
	ZipCode string `gorm:"not null" json:"zipcode" validate:"required"`
	Country string `gorm:"not null" json:"country" validate:"required"`
}

type GetOrderDetails struct {
	gorm.Model
	UserID  uint   `json:"user_id" validate:"required"`
	OrderID uint   `json:"order_id" validate:"required"`
	Street  string `gorm:"not null" json:"street" validate:"required"`
	City    string `gorm:"not null" json:"city" validate:"required"`
	State   string `gorm:"not null" json:"state" validate:"required"`
	ZipCode string `gorm:"not null" json:"zipcode" validate:"required"`
	Country string `gorm:"not null" json:"country" validate:"required"`
	Method  string `gorm:"not null" json:"method" vaidate:"required"`
}
