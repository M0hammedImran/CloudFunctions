package utils

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/m0hammedimran/CloudFunctions/EnterpriseRedirection/types"
	"github.com/mitchellh/mapstructure"
)

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

func VerifyJWT(authToken string, secretKey string) (*types.StandardClaims, error) {
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodECDSA)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	// parsing errors result
	if err != nil {
		return nil, err
	}

	var standardClaims = types.StandardClaims{}
	claims, ok := token.Claims.(jwt.MapClaims)
	mapstructure.Decode(claims, &standardClaims)

	// if there's a token
	if ok && token.Valid {
		return &standardClaims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func DecodeJwtClaim(authToken string) (types.StandardClaims, error) {
	var tokenSegments = strings.Split(authToken, ".")
	claimsBytes, err := jwt.DecodeSegment(tokenSegments[1])
	if err != nil {
		return types.StandardClaims{}, err
	}

	var standardClaims = types.StandardClaims{}
	err = json.Unmarshal(claimsBytes, &standardClaims)
	if err != nil {
		return types.StandardClaims{}, err
	}

	return standardClaims, nil
}
