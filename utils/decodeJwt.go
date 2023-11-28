package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/m0hammedimran/CloudFunctions/EnterpriseRedirection/types"
)

func decodeTokenSegment(segment string) ([]byte, error) {
	if l := len(segment) % 4; l > 0 {
		segment += strings.Repeat("=", 4-l)
	}
	return base64.URLEncoding.DecodeString(segment)
}

func DecodeAccessToken(accessToken string, claims types.StandardClaims) (*types.Token, error) {
	var err error
	token := &types.Token{Raw: accessToken}
	parts := strings.Split(accessToken, ".")
	var headerBytes []byte

	if headerBytes, err = decodeTokenSegment(parts[0]); err != nil {
		return token, fmt.Errorf("an error occurred decoding token segment")
	}

	if err = json.Unmarshal(headerBytes, &token.Header); err != nil {
		return token, fmt.Errorf("an error occurred unmarshalling header")
	}

	var claimBytes []byte
	if claimBytes, err = decodeTokenSegment(parts[1]); err != nil {
		return token, fmt.Errorf("an error occurred unmarshalling claims")
	}

	decoder := json.NewDecoder(bytes.NewBuffer(claimBytes))
	err = decoder.Decode(&claims)
	token.Claims = claims
	return token, err
}
