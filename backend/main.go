package main

import (
	"fmt"

	"github.com/fathima-sithara/ecommerce/user-service/config"
)

func init() {
	config.LoadEnvVariable()
	config.ConnectToDB()
	config.Pooling()
	config.AutoMigrate()
}

func main() {
	fmt.Println("welcoe to ecommerce")
}
