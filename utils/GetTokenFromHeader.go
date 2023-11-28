package utils

import (
	"fmt"
	"net/http"
	"strings"
)

func GetTokenFromHeader(header http.Header) (string, error) {
	auth := header.Get("Authorization")
	if auth == "" {
		return "", fmt.Errorf("no Authorization header found")
	}
	auth = strings.Replace(auth, "Bearer ", "", 1)

	if auth == "" {
		return "", fmt.Errorf("no Authorization header found")
	}
	return auth, nil
}
