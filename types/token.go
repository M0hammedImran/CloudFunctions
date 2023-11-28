package types

type Token struct {
	Raw       string                 `json:"raw"`       // The raw token.  Populated when you Parse a token
	Header    map[string]interface{} `json:"headers"`   // The first segment of the token
	Claims    StandardClaims         `json:"claims"`    // The second segment of the token
	Signature string                 `json:"signature"` // The third segment of the token.  Populated when you Parse a token
	Valid     bool                   `json:"valid"`     // Is the token valid?  Populated when you Parse/Verify a token
}

type StandardClaims struct {
	Id        int64 `json:"id,omitempty"`
	ExpiresAt int64 `json:"exp,omitempty"`
	IssuedAt  int64 `json:"iat,omitempty"`
}

type MapClaims map[string]interface{}
