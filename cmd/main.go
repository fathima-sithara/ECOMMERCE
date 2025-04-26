package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/fathimasithara01/ecommerce/database"
	"github.com/fathimasithara01/ecommerce/routes"
)

func main() {
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "8080"
	}

	database.InitDB()

	router := gin.Default()
	// router.LoadHTMLGlob("templates/*")
	routes.UserRoutes(router)
	routes.AdminRoutes(router)

	router.Run()
}
