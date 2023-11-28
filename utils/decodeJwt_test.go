package utils

import (
	"testing"

	"github.com/m0hammedimran/CloudFunctions/EnterpriseRedirection/types"
)

func TestDecodeJwt(t *testing.T) {
	var token = ""

	_, err := DecodeAccessToken(token, types.StandardClaims{})
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTc1NCwiaWF0IjoxNzAwNzMxNDQwLCJleHAiOjE3MDMzMjM0NDB9.usP-o72fvOcCOMc5cr34_szXD0b6Z0UQ_MiHeTavm7I"
	decodedToken, err := DecodeAccessToken(token, types.StandardClaims{})
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if decodedToken.Claims.Id != 1754 {
		t.Errorf("Expected 1754, got %v", decodedToken.Claims.Id)
	}
}
