package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/fathimasithara01/ecommerce/models"
)

var Db *gorm.DB

func InitDB() *gorm.DB {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}
	Db = connectDB()

	return Db
}

func connectDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Connected to database")

	// Migrate all models at once
	err = db.AutoMigrate(
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

	return db
}
