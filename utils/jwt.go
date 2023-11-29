package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	UserId int `json:"id,omitempty"`
	jwt.RegisteredClaims
}

func GenerateJWT(secretKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = "username"
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "Signing Error", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string, secretKey string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*MyClaims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("unknown claims type, cannot proceed")
}
