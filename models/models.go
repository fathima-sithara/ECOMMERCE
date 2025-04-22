package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey;unique"`
	First_Name   string `json:"first_name" gorm:"not null" validate:"required,miin=2,max=50"`
	Last_Name    string `json:"last_name" gorm:"not null" validate:"required,min=2,max=50"`
	Email        string `json:"email" gorm:"not null;unique" validate:"eamil,required"`
	Password     string `json:"password" gorm:"not null" validate:"required"`
	Phone        string `json:"phone" gorm:"not null;unique" validate:"required"`
	Otp          string `json:"otp"`
	Block_Status bool   `json:"block_status" gorm:"not null"`
	Verified     bool   `json:"verified" gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Admin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Product struct {
	gorm.Model
	ProductId   uint     `json:"product_id" gorm:"autoIncrement"`
	ProductName string   `json:"product_name" gorm:"not null"`
	Price       string   `json:"price" gorm:"not null"`
	Image       string   `json:"image" gorm:"not null"`
	Stock       uint     `json:"stock"`
	Color       string   `json:"color" gorm:"not null"`
	Description string   `json:"description"`
	Brand       string   `gorm:"ForeignKey:BrandId"`
	BrandId     uint     `json:"brand_id"`
	Categery    Categery `gorm:"ForeignKey:CatogreyId"`
	CatogreyId  uint     `json:"Catogrey_id"`
}
type Brand struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Brands string `json:"category_id" gorm:"autoIncrement"`
}

type Categery struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	CategoryID uint   `json:"json" gorm:"autoIncrement"`
	Categery   string `json:"categery"`
}

type Cart struct {
	gorm.Model
	Product    Product `gorm:"ForeignKey:ProductId"`
	PorductID  uint
	Quantity   uint
	Price      uint
	TotalPrice uint
	Userid     uint
	User       User `gorm:"ForeignKey:Userid"`
}

type Address struct {
	AddressID uint `JSON:"addressid" gorm:"primaryKey;unique"`
	User      User `gorm:"ForeignKey:Userid"`
	Userid    uint `JSON:"uid"`

	Name       string `JSON:"name" gorm:"not null"`
	Phoneno    string `JSON:"phoneno" gorm:"not null"`
	Houseno    string `JSON:"houseno" gorm:"not null"`
	Area       string `JSON:"area" gorm:"not null"`
	Landmark   string `JSON:"landmark" gorm:"not null"`
	City       string `JSON:"city" gorm:"not null"`
	Pincode    string `JSON:"pincode" gorm:"not null"`
	District   string `JSON:"district" gorm:"not null"`
	State      string `JSON:"state" gorm:"not null"`
	Country    string `JSON:"country" gorm:"not null"`
	Defaultadd bool   `JSON:"defaultadd" gorm:"default:false"`
}

type Wallet struct {
	Id     uint
	User   User `gorm:"ForeignKey:UserId"`
	UserId uint
	Amount float64
}

type WalletHistory struct {
	Id             uint `JSON:"Id" gorm:"primarykey"`
	User           User `gorm:"ForeignKey:UserId"`
	UserId         uint
	Amount         float64
	TransctionType string
	Date           time.Time
}

type Coupon struct {
	ID            int
	CouponCode    string
	DiscountPrice float64
	CreatedAt     time.Time
	Expired       time.Time
}

type Wishlist struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	Userid    uint
	User      User    `gorm:"ForeignKey:Userid"`
	Product   Product `gorm:"ForeignKey:ProductId"`
	ProductId uint
}
