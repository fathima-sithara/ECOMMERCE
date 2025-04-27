package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string `json:"first_name" gorm:"not null" validate:"required,min=2,max=50"`
	LastName    string `json:"last_name" gorm:"not null" validate:"required,min=2,max=50"`
	Email       string `json:"email" gorm:"not null;unique" validate:"required"`
	Password    string `json:"password" gorm:"not null" validate:"required"`
	Phone       string `json:"phone" gorm:"not null;" `
	Otp         string `json:"otp"`
	BlockStatus bool   `json:"block_status" gorm:"not null;default:false"`
	Verified    bool   `json:"verified" gorm:"not null;default:false"`
}

type Admin struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Product struct {
	gorm.Model
	ProductName string   `json:"product_name" gorm:"not null"`
	Price       string   `json:"price" gorm:"not null"`
	Image       string   `json:"image" gorm:"not null"`
	Stock       uint     `json:"stock"`
	Color       string   `json:"color" gorm:"not null"`
	Description string   `json:"description"`
	BrandId     uint     `json:"brand_id"`
	Brand       Brand    `gorm:"foreignKey:BrandId"`
	CatogreyId  uint     `json:"catogrey_id"`
	Categery    Categery `gorm:"foreignKey:CatogreyId"`
}

type Brand struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Brands string `json:"brands"`
}

type Categery struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	CategoryID uint   `json:"category_id" gorm:"autoIncrement"`
	Categery   string `json:"categery"`
}

type Cart struct {
	gorm.Model
	ProductId  uint    `json:"product_id"`
	Product    Product `gorm:"foreignKey:ProductId"`
	Quantity   uint    `json:"quantity"`
	Price      uint    `json:"price"`
	TotalPrice uint    `json:"total_price"`
	UserId     uint    `json:"user_id"`
	User       User    `gorm:"foreignKey:UserId"`
}

type Address struct {
	AddressID  uint   `json:"addressid" gorm:"primaryKey;unique"`
	UserId     uint   `json:"uid"`
	User       User   `gorm:"foreignKey:UserId"`
	Name       string `json:"name" gorm:"not null"`
	Phoneno    string `json:"phoneno" gorm:"not null"`
	Houseno    string `json:"houseno" gorm:"not null"`
	Area       string `json:"area" gorm:"not null"`
	Landmark   string `json:"landmark" gorm:"not null"`
	City       string `json:"city" gorm:"not null"`
	Pincode    string `json:"pincode" gorm:"not null"`
	District   string `json:"district" gorm:"not null"`
	State      string `json:"state" gorm:"not null"`
	Country    string `json:"country" gorm:"not null"`
	Defaultadd bool   `json:"defaultadd" gorm:"default:false"`
}

type Wallet struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	UserId uint    `json:"user_id"`
	User   User    `gorm:"foreignKey:UserId"`
	Amount float64 `json:"amount"`
}

type WalletHistory struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	UserId         uint      `json:"user_id"`
	User           User      `gorm:"foreignKey:UserId"`
	Amount         float64   `json:"amount"`
	TransctionType string    `json:"transaction_type"`
	Date           time.Time `json:"date"`
}

type Coupon struct {
	ID            int       `json:"id"`
	CouponCode    string    `json:"coupon_code"`
	DiscountPrice float64   `json:"discount_price"`
	CreatedAt     time.Time `json:"created_at"`
	Expired       time.Time `json:"expired"`
}

type Wishlist struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	UserId    uint    `json:"user_id"`
	User      User    `gorm:"foreignKey:UserId"`
	ProductId uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductId"`
}
