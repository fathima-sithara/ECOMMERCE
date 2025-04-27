package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func GenerateJWT(userID uint, email string, role string) (string, error) {
	expiration, err := time.ParseDuration(os.Getenv("jwt.expiration"))
	if err != nil {
		return "", err
	}
	// fmt.Println("lsfkjglsfjg", viper.GetString("jwt.expiration"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    userID,
		"email": email,
		"role":  role, // Include role in the JWT token
		"exp":   time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func ExtractClaims(tokenString string) (uint, string, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("jwt.secret")), nil // Use correct secret key
	})

	if err != nil {
		return 0, "", "", fmt.Errorf("invalid token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, okID := claims["id"].(float64) // JWT stores numbers as float64
		email, okEmail := claims["email"].(string)
		role, okRole := claims["role"].(string)

		if !okID || !okEmail || !okRole {
			return 0, "", "", fmt.Errorf("invalid claim format")
		}

		return uint(id), email, role, nil
	}

	return 0, "", "", fmt.Errorf("invalid token claims")
}
