package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name" gorm:"not null" validate:"required,min=2,max=100"`
	Description string         `json:"description" gorm:"not null" validate:"required,min=10,max=500"`
	Price       float64        `json:"price" gorm:"not null" validate:"required,min=0"`
	Stock       int            `json:"stock" gorm:"not null" validate:"required,min=0"`
	IsAvailable bool           `json:"is_available" gorm:"default:true"`
	CompanyName string         `json:"company_name" validate:"omitempty,min=2,max=100"`
	Brand       string         `json:"brand" validate:"omitempty,min=2,max=50"`
	Size        pq.StringArray `gorm:"type:text[]" json:"size" validate:"dive,required"`
	Category    string         `json:"category" validate:"omitempty,min=2,max=50"`
	Images      []ProductImage `gorm:"foreignKey:ProductID" json:"images"`
	Reviews     []Review       `gorm:"foreignKey:ProductID" json:"reviews"`
}

type ProductImage struct {
	gorm.Model
	ProductID uint   `json:"product_id" gorm:"not null" validate:"required"`
	URL       string `json:"url" gorm:"not null" validate:"required,url"`
	IsMain    bool   `json:"is_main" gorm:"default:false"`
}

type Filter struct {
	MinPrice    *float64 `form:"min_price"`
	MaxPrice    *float64 `form:"max_price"`
	IsAvailable *bool    `form:"is_available"`
	Category    string   `form:"category"`
	Brand       string   `form:"brand"`
}
