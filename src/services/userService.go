package services

import (
	"fmt"
	"log"

	"github.com/fathimasithara01/ecommerce/helpers"
	"github.com/fathimasithara01/ecommerce/src/models"
	"github.com/fathimasithara01/ecommerce/src/repository"
)

type UserServices struct {
}

func (us *UserServices) RegisterUser(req models.User) error {
	hash, err := helpers.HashPassword(req.Password)
	if err != nil {
		return err
	}
	// Log IPgSQLrepo state to check if it's nil
	if repository.IPgSQLrepo == nil {
		log.Println("ERROR: IPgSQLrepo is nil!")
		return fmt.Errorf("repository is nil")
	}

	req.Password = hash
	return repository.IPgSQLrepo.Insert(&req)
}
