package jwt

import (
	"github.com/golang-jwt/jwt"
)

type TokenValidator struct {
	secretKey string
}

func NewTokenValidator(cfg Config) *TokenValidator {
	return &TokenValidator{
		secretKey: cfg.SecretKey,
	}
}

type TokenParsedParams struct {
	Id int64
}

func (v *TokenValidator) ValidateToken(token string) (*TokenParsedParams, error) {
	claims := &Claims{}
	parsed, err := jwt.ParseWithClaims(token, claims, func(jwtToken *jwt.Token) (interface{}, error) {
		return []byte(v.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !parsed.Valid {
		return nil, err
	}

	return &TokenParsedParams{Id: claims.Id}, nil
}
