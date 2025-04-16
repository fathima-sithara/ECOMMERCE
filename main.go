package main

import (
	"fmt"

	"github.com/fathima-sithara/ecommerce/initalizers"
	// "github.com/gin-gonic/gin"
)

func init() {
	initalizers.LoadEnvVariable()
	initalizers.ConnectToDB()
	initalizers.Pooling()
	initalizers.AutoMigrate()
}

func main() {
	// r:= gin.Default{}
	fmt.Println("welcomme to skinglow ecommerce")
}
