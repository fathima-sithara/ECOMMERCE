package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string         `gorm:"not null" json:"name" validate:"required,min=3,max=100"`
	Description  string         `gorm:"not null" json:"description" validate:"required,min=10,max=100"`
	Price        string         `gorm:"not null" json:"price" validate:"required,min=0"`
	Stock        int            `gorm:"not null" json:"stock" validate:"required,min=0"`
	IsAvailable  bool           `gorm:"default:true" json:"is_available"`
	ComapanyName string         `json:"company_name" validate:"omitempty,min=2,max=100"`
	Brand        string         `json:"brand" validate:"omitempty,min=3,max=100"`
	Size         pq.StringArray `gorm:"type:text[]" json:"size" validate:"dive,required"`
	Category     string         `json:"category" validate:"omitempty,min=2,max=50"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	Images       []ProductImage `gorm:"foreignKey:ProductID" json:"images"`
	Reviews      []Review       `gorm:"foreignKey:ProductID" json:"reviews"`
}

type ProductImage struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `json:"product_id" gorm:"not null" validate:"required"`
	URL       string    `json:"url" gorm:"not null" validate:"required,url"`
	IsMain    bool      `json:"is_main" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Filter struct {
	MinPrice    *float64 `form:"min_price"`
	MaxPrice    *float64 `form:"max_price"`
	IsAvailable *bool    `form:"is_available"`
	Category    string   `form:"category"`
	Brand       string   `form:"brand"`
}
