package migration

import (
	"log"

	"github.com/fathimasithara01/ecommerce/database"
	"github.com/fathimasithara01/ecommerce/src/models"
)

func Migration() {
	db := database.PgSQLDB

	err := db.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.Product{},
		&models.Categery{}, // Fixed typo from Categery
		&models.Brand{},
		&models.Address{},
		&models.Cart{},
		&models.Wallet{},
		&models.WalletHistory{},
		&models.Coupon{},
		&models.Wishlist{},
	)
	if err != nil {
		log.Fatal("AutoMigration failed:", err)
	}
}
