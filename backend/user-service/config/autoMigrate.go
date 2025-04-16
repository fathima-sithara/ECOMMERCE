package config

import "github.com/fathima-sithara/ecommerce/model"

// Config files (env, DB config, etc.)

func AutoMigrate() {
	DB.AutoMigrate(
		&model.User{},
		&model.Address{},
		&model.Cart{},
		&model.CartItem{},
		&model.Order{},
		&model.GetOrderDetils{},
		&model.Product{},
		&model.Review{},
		&model.Wishlist{},
		&model.WishlistItem{},
		&model.OrderItem{},
		&model.Admin{},
		&model.UserToken{},
		&model.ProductImage{},
		&model.Payment{},
	)
}
