package models

import "gorm.io/gorm"

type UserToken struct {
	gorm.Model
	UserID uint   `gorm:"index" json:"userid"`
	Token  string `gorm:"type varchar(255);not null" json:"token"`
}

type User struct {
	gorm.Model
	FirstName         string      `json:"firstname" validate:"required,min=2,max=100"`
	LastName          string      `json:"lastname" validate:"required,min=3,max=100"`
	Email             string      `gorm:"unique,type:varchar(100)" json:"eamil" validate:"required,email"`
	Password          string      `json:"password" validate:"required,min=10,max=50"`
	Phone             string      `json:"phone"  validate:"required,min=10"`
	Ban               bool        `json:"ban"`
	ProfilePictureURL string      `gorm:"type:varchar(255)" json:"profile_picture_url"`
	IsVerified        bool        `gorm:"default:false" json:"is_verified"`
	Cart              Cart        `gorm:"foreignKey:UserID" json:"cart"`
	Orders            []order     `gorm:"foreignKey:UserID" json:"orders"`
	Address           []Address   `gorm:"foreignKey:UserID" json:"address"`
	Wishlist          Wishlist    `gorm:"foreignKey:UserID" json:"wishlist"`
	Reviews           []Review    `gorm:"foreignKey:UserID" json:"reviews"`
	Token             []UserToken `gorm:"foreignKey:UserID" json:"token"`
}
