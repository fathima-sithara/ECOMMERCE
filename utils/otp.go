package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/smtp"
	"os"
	"strconv"
)

// GenerateOTP creates a random 4-digit OTP
func GenerateOTP() (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(9000))
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(n.Int64()+1000, 10), nil
}

// SendOTPEmail sends the OTP to the provided email address
func SendOTPEmail(toEmail, otp string) error {
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	body := fmt.Sprintf("Subject: OTP Verification\n\nYour OTP is: %s", otp)

	auth := smtp.PlainAuth("", from, password, smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, []byte(body))
}

// VerifyAndSendOTP handles OTP generation and sending
func VerifyAndSendOTP(email string) (string, error) {
	otp, err := GenerateOTP()
	if err != nil {
		return "", fmt.Errorf("failed to generate OTP: %w", err)
	}
	if err := SendOTPEmail(email, otp); err != nil {
		return "", fmt.Errorf("failed to send email: %w", err)
	}
	return otp, nil
}
