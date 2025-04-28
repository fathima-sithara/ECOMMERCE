package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/fathimasithara01/ecommerce/config"
	"github.com/fathimasithara01/ecommerce/database"
	"github.com/fathimasithara01/ecommerce/migration"
	"github.com/fathimasithara01/ecommerce/routes"
	"github.com/fathimasithara01/ecommerce/src/repository"
	validator "github.com/fathimasithara01/ecommerce/utils/validator"
)

func main() {
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "8080"
	}

	validator.Init()
	config.LoadConfig()
	database.GetInstancePostgres()
	migration.Migration()
	repository.PgSQLInit()
	router := gin.Default()
	// router.LoadHTMLGlob("templates/*")

	routes.Routes(router)
	router.Run()
}
