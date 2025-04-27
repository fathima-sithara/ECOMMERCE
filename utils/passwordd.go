package utils

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/fathimasithara01/ecommerce/src/models"
)

var admin models.Admin
var user models.User

// for admin side
func AdminHashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	admin.Password = string(bytes)
	return nil
}

func AdminCheckPassword(inputPassword string) error {

	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(inputPassword))
	if err != nil {
		return err
	}
	return nil
}

// for user side
func UserHashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func UserCheckPassword(inputPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputPassword))
	if err != nil {
		return err
	}
	return nil
}
