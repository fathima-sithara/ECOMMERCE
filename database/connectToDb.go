package database

import (
	"fmt"
	"log"
	"os"

	"github.com/fathima-sithara/ecommerce/models"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() *gorm.DB {
	Db = connectDB()

	return Db
}

func connectDB() *gorm.DB {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf("host =%s user= %s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed tto connect database", err)
	}
	fmt.Println("/n connected to database:", db.Name())

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Admin{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Categery{})
	db.AutoMigrate(&models.Brand{})
	db.AutoMigrate(&models.Address{})
	db.AutoMigrate(&models.Categery{})
	db.AutoMigrate(&models.Cart{})
	db.AutoMigrate(&models.Address{})
	db.AutoMigrate(&models.Wallet{})
	db.AutoMigrate(&models.WalletHistory{})
	db.AutoMigrate(&models.Coupon{})
	db.AutoMigrate(&models.Wishlist{})

	return db
}
