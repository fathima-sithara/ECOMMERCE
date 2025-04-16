package initalizers

import "github.com/fathima-sithara/ecommerce/models"

func AutoMigrate() {
	DB.AutoMigrate(
		&models.User{},
		&models.Address{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.GetOrderdetils{},
		&models.Product{},
		&models.Review{},
		&models.Wishlist{},
		&models.WishlistItem{},
		&models.OrderItem{},
		&models.Admin{},
		&models.UserToken{},
		&models.ProductImage{},
		&models.Payment{},
	)
}
