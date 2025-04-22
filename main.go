package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/fathima-sithara/ecommerce/database"
	"github.com/fathima-sithara/ecommerce/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	database.InitDB()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	routes.UserRoutes(router)
	routes.AdminRoutes(router)

	router.Run()
}
