package jwtToken

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
	jwt.StandardClaims
}

func GenerateJWT(cpf string, secret string) (string, error) {
	JWT_KEY := []byte(os.Getenv("JWT_KEY"))
	expirationTime := time.Now().Add(1 * time.Minute)
	claims := &JWTClaim{
		CPF:    cpf,
		Secret: secret,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWT_KEY)
	return tokenString, err
}

func ValidateToken(signedToken string) error {
	JWT_KEY := []byte(os.Getenv("JWT_KEY"))
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_KEY), nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return errors.New("couldn't parse claims")

	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return errors.New("token expired")
	}
	return nil
}

func GetDocument(tokenString string) (string, error) {
	JWT_KEY := []byte(os.Getenv("JWT_KEY"))
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JWT_KEY, nil
	})

	if err != nil {
		return "", err
	}

	// do something with decoded claims
	for key, val := range claims {
		if key == "cpf" {
			return string(fmt.Sprintf("%v", val)), nil
		}
	}

	return "", errors.New("JWT key not found")
}
