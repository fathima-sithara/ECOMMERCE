package model

import (
	"time"

	"gorm.io/gorm"
)

type UserToken struct {
	gorm.Model
	UserID uint   `gorm:"index" json:"user_id"`
	Token  string `gorm:"type:varchar(255);not null" json:"token"`
}

type User struct {
	ID                uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Email             string         `gorm:"unique;type:varchar(100)" json:"email" validate:"required,email"`
	FirstName         string         `json:"first_name" validate:"required,min=2,max=100"`
	LastName          string         `json:"last_name" validate:"required,min=2,max=100"`
	Password          string         `json:"password" validate:"required,min=8,max=255"`
	Phone             string         `json:"phone" validate:"required,min=10"`
	Ban               bool           `json:"ban"`
	ProfilePictureURL string         `gorm:"type:varchar(255)" json:"profile_picture_url,omitempty"`
	IsVerified        bool           `gorm:"default:false" json:"is_verified"`
	IsActive          bool           `gorm:"default:true" json:"is_active"`
	CreatedAt         time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Cart              Cart           `gorm:"foreignKey:UserID" json:"cart"`
	Orders            []Order        `gorm:"foreignKey:UserID" json:"orders"`
	Addresses         []Address      `gorm:"foreignKey:UserID" json:"addresses"`
	Wishlist          Wishlist       `gorm:"foreignKey:UserID" json:"wishlist"`
	Reviews           []Review       `gorm:"foreignKey:UserID" json:"reviews"`
	Token             []UserToken    `gorm:"foreignKey:UserID" json:"spurt_tokens,omitempty"`
}
