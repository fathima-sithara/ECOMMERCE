package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type EmailSender struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type EmailRecipient struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type EmailContent struct {
	Sender  EmailSender      `json:"sender"`
	To      []EmailRecipient `json:"to"`
	Subject string           `json:"htmlContent"`
}

func SendEmail(toEmail, toName, subject, body string) error {
	email := EmailContent{
		Sender: EmailSender{
			Name:  viper.GetString("brevo.sender_name"),
			Email: viper.GetString("brevo.sender_email"),
		},
		To: []EmailRecipient{
			{
				Email: toEmail,
				Name:  toName,
			},
		},
		Subject:     subject,
		// HtmlContent: body,
	}

	emailData, err := json.Marshal(email)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://api.brevo.com/v3/smtp/email", bytes.NewBuffer(emailData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("api-key", viper.GetString("brevo.api_key"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to send email: %s", resp.Status)
	}

	return nil
}

// Generate a random 4-digit OTP
func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%04d", rand.Intn(10000))
}
