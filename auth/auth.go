package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte(os.Getenv("SECRET"))

type JWTClaim struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

// GenerateJWT creates both access and refresh tokens
func GenerateJWT(id string) (map[string]string, error) {
	// Access Token Claims
	expirationTime := time.Now().Add(1 * time.Hour)
	accesClaims := &JWTClaim{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			// ExpiresAt: expirationTime,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   id,
		},
	}

	// Create access token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, accesClaims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return nil, err
	}

	// Refresh Token Claims (simpler, using MapClaims)
	refrehToken := jwt.New(jwt.SigningMethodES256)      // Creating a New JWT (Refresh Token)
	refreshClaims := refrehToken.Claims.(jwt.MapClaims) //Setting the Claims (Data) for the Token
	refreshClaims["id"] = id
	refreshClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign refresh token (ideally use a separate secret for refresh tokens)
	reefreshTokenString, err := refrehToken.SignedString([]byte("secret")) // Sign the token using a secret key to ensure it’s secure and Return the signed refresh token.
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  tokenString,
		"refresh_Token": reefreshTokenString,
	}, nil
}

var P string //// Global variable to store the ID from the claims

// ValidateToken validates the JWT token
func ValidateToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		},
	)

	if err != nil || !token.Valid {
		return errors.New("invalid token")
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return errors.New("couldn't parse claims")
	}

	// Set the global user ID from claims for use in the middleware
	P = claims.Id

	// Check token expiration
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return errors.New("token expired")
	}

	return nil
}
