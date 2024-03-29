package token

import (
	"crypto/rsa"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// JWTTokenVerifier verifies jwt access tokens.
type JWTTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

// Verify verifies a token and returns account id.
func (v *JWTTokenVerifier) Verify(token string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(*jwt.Token) (interface{}, error) {
		return v.PublicKey, nil
	})

	if err != nil {
		return "", fmt.Errorf("cannot parse token: %v", err)
	}

	if !t.Valid {
		return "", fmt.Errorf("token not valid")
	}

	clm, ok := t.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return "", fmt.Errorf("token claim is not StandardClaims")
	}

	return clm.Subject, nil
}
