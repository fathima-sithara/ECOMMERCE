// emial-based otp(using gmail SMTP)

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
	from := os.Getenv("EMAIL")        // Your Gmail email address
	password := os.Getenv("PASSWORD") // Your Gmail app password (not your Google password)
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

//phone -based OTP(using Twilio)

// package utils

// import (
// 	"crypto/rand"
// 	"fmt"
// 	"math/big"
// 	"os"
// 	"strconv"
//

// 	"github.com/twilio/twilio-go"
// 	openapi "github.com/twilio/twilio-go/rest/api/v2010"
// )

// // GenerateOTP creates a 6-digit random OTP
// func GenerateOTP(length int) (string, error) {
// 	var otp string
// 	for i := 0; i < length; i++ {
// 		num, err := rand.Int(rand.Reader, big.NewInt(10))
// 		if err != nil {
// 			return "", err
// 		}
// 		otp += strconv.FormatInt(num.Int64(), 10)
// 	}
// 	return otp, nil
// }

// // SendOTPViaTwilio sends OTP using Twilio SMS
// func SendOTPViaTwilio(phoneNumber, otp string) error {
// 	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
// 	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
// 	fromPhone := os.Getenv("TWILIO_PHONE_NUMBER")

// 	client := twilio.NewRestClientWithParams(twilio.ClientParams{
// 		Username: accountSid,
// 		Password: authToken,
// 	})

// 	message := fmt.Sprintf("Your OTP code is: %s. Do not share this with anyone.", otp)

// 	params := &openapi.CreateMessageParams{}
// 	params.SetTo(phoneNumber)
// 	params.SetFrom(fromPhone)
// 	params.SetBody(message)

// 	_, err := client.Api.CreateMessage(params)
// 	if err != nil {
// 		return fmt.Errorf("failed to send OTP via Twilio: %w", err)
// 	}

// 	return nil
// }

// // VerifyAndSendOTP handles OTP generation and sending via Twilio
// func VerifyAndSendOTP(phoneNumber string) (string, error) {
// 	otp, err := GenerateOTP(6)
// 	if err != nil {
// 		return "", err
// 	}
// 	if err := SendOTPViaTwilio(phoneNumber, otp); err != nil {
// 		return "", err
// 	}
// 	return otp, nil
// }
