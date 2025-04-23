package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/fathima-sithara/ecommerce/config"
	"github.com/fathima-sithara/ecommerce/routes"
)

func main() {
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "8080"
	}

	config.InitDB()

	router := gin.Default()
	// router.LoadHTMLGlob("templates/*")
	routes.UserRoutes(router)
	routes.AdminRoutes(router)

	router.Run()
}
