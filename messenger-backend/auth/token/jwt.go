package token

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtTokenGen struct {
	privateKey *rsa.PrivateKey
	issuer     string
	nowFunc    func() time.Time
}

func NewJwtTokenGen(issuer string, privateKey *rsa.PrivateKey) *JwtTokenGen {
	return &JwtTokenGen{
		privateKey: privateKey,
		issuer:     issuer,
		nowFunc:    time.Now,
	}
}

func (j *JwtTokenGen) GenerateToken(uid string, expire time.Duration) (token string, err error) {
	now := j.nowFunc()
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.RegisteredClaims{
		Issuer:    j.issuer,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(expire)),
		Subject:   uid,
	})
	return tkn.SignedString(j.privateKey)
}
