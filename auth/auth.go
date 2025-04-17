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
	jwt.StandardClaims
}

func GenerateJWT(id string) (map[string]string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return nil, err
	}

	refrehToken := jwt.New(jwt.SigningMethodES256)
	rTokenClaims := refrehToken.Claims.(jwt.MapClaims)
	rTokenClaims["id"] = id

	rTokenClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	rt, err := refrehToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  tokenString,
		"refresh_Token": rt,
	}, nil
}

var P string

func VlidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	P = claims.Id
	if !ok {
		err = errors.New("couldn't parsw claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
