package helpers

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Verify password
func CheckPasswordHash(hash, password string) bool {
	//cheoing password
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

// Generate a 6-digit OTP
func GenerateOTP() string {
	n, err := rand.Int(rand.Reader, big.NewInt(900000)) // Range: 0 to 899999
	if err != nil {
		log.Fatal(err)
	}
	otp := 100000 + n.Int64() // Shift range to 100000 - 999999
	return fmt.Sprintf("%06d", otp)
}

func StringToInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid integer: %v", err)
	}
	return num, nil
}

func IntToString(n int) string {
	return strconv.Itoa(n)
}
