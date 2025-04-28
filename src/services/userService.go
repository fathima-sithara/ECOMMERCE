package services

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/fathimasithara01/ecommerce/database"
	"github.com/fathimasithara01/ecommerce/helpers"
	"github.com/fathimasithara01/ecommerce/src/models"
	"github.com/fathimasithara01/ecommerce/src/repository"
	"github.com/fathimasithara01/ecommerce/utils/jwt"
	"github.com/fathimasithara01/ecommerce/utils/utils"
)

type UserServices struct{}

func (s *UserServices) db() *gorm.DB {
	return database.PgSQLDB
}

func (s *UserServices) Register(username, email, password, phone string) (string, error) {
	var existingUser models.User
	//check email already exist
	if err := repository.IPgSQLRepo.FindByEmail(&existingUser, email); err == nil {
		return "", fmt.Errorf("email already exist")
	}
	// check if username already exists
	if err := s.db().Where("username=?", username).First(&existingUser).Error; err != nil {
		return "", fmt.Errorf("username alreaddy exist")
	}

	// create a new user
	var newUser models.User
	newUser.UserName = username
	newUser.Email = email
	newUser.Password = password
	newUser.Phone = phone

	//  Hash the password
	hash, err := helpers.HashPassword(password)
	if err != nil {
		return "", err
	}
	newUser.Password = hash

	//generate OTP and send verification email
	subject := "Email Verification OTP"
	otp := helpers.GenerateOTP()
	body := fmt.Sprintf("Hello %s <br> This is your OTP for email verification :%s", username, otp)
	if err := utils.SendEmail(newUser.UserName, newUser.Email, subject, body); err != nil {
		return "", err
	}
	newUser.OtpCreatedAt = time.Now()
	newUser.Otp = otp
	if err := repository.IPgSQLRepo.Insert(&newUser); err != nil {
		return "", err
	}
	otpInt, err := helpers.StringToInt(otp)
	if err != nil {
		return "", err
	}
	tokenString, err := jwt.GenerateJWT(uint(otpInt), newUser.Email, "email verification")
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *UserServices) Login(email, password string) (string, models.User, bool, error) {
	var user models.User

	err := repository.IPgSQLRepo.FindByEmail(&user, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", user, false, errors.New("We couldn't find any account associated with this email. Please verify your credentials or register.")
		}
		return "", user, false, err
	}
	if user.Verified == false {
		subject := "Email Verification OTP"
		otp := helpers.GenerateOTP()
		user.Otp = otp
		user.OtpCreatedAt = time.Now()
		body := fmt.Sprintf("Hello %s,This is your OTP for email verification: %s", user.UserName, otp)
		if err := utils.SendEmail(user.Email, user.UserName, subject, body); err != nil {
			return "", user, false, err
		}
		err := repository.IPgSQLRepo.Save(&user)
		if err != nil {
			return "", user, false, err
		}
		return "", user, true, errors.New("pleas verify user email address.")
	}

	check := helpers.CheckPasswordHash(user.Password, password)
	fmt.Println(check, user.Password, ">>>", password)
	if !check {
		return "", user, false, errors.New("The password you entered is incorrect. Please try again.")
	}

	tokenString, err := jwt.GenerateJWT(user.ID, user.Email, "user")
	if err != nil {
		return "", user, false, err
	}
	return tokenString, user, true, nil
}
func (s *UserServices) OtpService(email, meth string) (string, string, error) {
	var user models.User
	if meth == "email_verification" || meth == "reset_password" {
		// Find user by email
		if err := repository.IPgSQLRepo.FindByEmail(&user, email); err != nil {
			return "", "", err
		}
		if meth == "reset_password" {
			if !user.Verified {
				return "", "", errors.New("pleas verify your email")
			}
		}
		// Check if OTP was generated recently
		timeSinceLastOtp := time.Since(user.OtpCreatedAt)
		if timeSinceLastOtp < 60*time.Second {
			remainingTime := 60 - int(timeSinceLastOtp.Seconds())
			return fmt.Sprintf("Please wait %d seconds before requesting a new OTP", remainingTime), "", nil
		}

		// Generate a new OTP
		otp := helpers.GenerateOTP()
		user.Otp = otp
		user.OtpCreatedAt = time.Now()

		// Prepare email
		subject := "Email Verification OTP"
		body := fmt.Sprintf(
			`<html><head></head><body><p>Hello %s,</p><p>This is your OTP for email verification: %s</p></body></html>`,
			user.UserName, otp,
		)

		// Send OTP via email
		if err := utils.SendEmail(user.Email, user.UserName, subject, body); err != nil {
			return "", "", err
		}

		// Save updated user OTP details
		if err := repository.IPgSQLRepo.Save(&user); err != nil {
			return "", "", err
		}
		otpInt, err := helpers.StringToInt(otp)
		if err != nil {
			return "", "", err
		}
		tokenString, err := jwt.GenerateJWT(uint(otpInt), user.Email, meth)
		if err != nil {
			return "", "", err
		}

		return "OTP sent successfully", tokenString, nil
	}

	return "", "", errors.New("your methoud not current pleas enter proper way..!")

}
func (s *UserServices) OtpVerifiecation(otp, token string) (string, error) {
	tokenOtp, email, meth, err := jwt.ExtractClaims(token)
	if err != nil {
		return "", err
	}
	reqStrOtp := helpers.IntToString(int(tokenOtp))

	if reqStrOtp != otp {
		return "", errors.New("pleas enter valid otp")
	}
	var u models.User
	if err := repository.IPgSQLRepo.FindByEmail(&u, email); err != nil {
		return "", errors.New("this email not restered")
	}
	if time.Since(u.OtpCreatedAt) > 3*time.Minute {
		return "", errors.New("OTP has expired, please request a new one")
	}
	if u.Otp != otp {
		return "", errors.New("otp not match")
	}
	if meth == "email_verification" {
		u.Verified = true
		u.Otp = ""
	} else if meth != "reset_password" {
		return "", errors.New("something went wrong")
	}

	if err = repository.IPgSQLRepo.Save(&u); err != nil {
		return "", err
	}
	if meth == "reset_password" {
		tokenString, err := jwt.GenerateJWT(tokenOtp, u.Email, "new_password")
		if err != nil {
			return "", err
		}
		return tokenString, nil
	}

	return "", nil
}

func (s *UserServices) ForgotPassword(pass, token string) error {
	tokenOtp, email, meth, err := jwt.ExtractClaims(token)
	if err != nil {
		return err
	}
	reqStrOtp := helpers.IntToString(int(tokenOtp))

	if meth != "new_password" {
		return errors.New("something went wrong")
	}
	var user models.User
	if err := repository.IPgSQLRepo.FindByEmail(&user, email); err != nil {
		return err
	}
	if user.Verified == false {
		return errors.New("pleas veify user")
	}
	if time.Since(user.OtpCreatedAt) > 3*time.Minute {
		return errors.New("OTP has expired, please request a new one")
	}
	if reqStrOtp != user.Otp {
		return errors.New("pleas enter valid otp")
	}
	user.Otp = ""
	hPas, err := helpers.HashPassword(pass)
	if err != nil {
		return err
	}

	user.Password = hPas
	if err := repository.IPgSQLRepo.Save(&user); err != nil {
		return err
	}
	return nil
}

func (s *UserServices) ValidateUserName(name string) error {
	var u models.User
	err := s.db().Where("user_name = ? ", name).First(&u).Error
	if err == nil || u.ID == 0 {
		return nil
	}
	return errors.New("this user name already in use")
}
