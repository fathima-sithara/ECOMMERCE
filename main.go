package main

import (
	"os"

	"github.com/fathima-sithara/ecommerce/database"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	database.InitDB()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Run()
}
