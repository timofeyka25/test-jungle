package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenGenerator struct {
	secretKey string
}

type Claims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

type Params struct {
	Id  int
	Ttl time.Duration
}

func NewTokenGenerator(cfg Config) *TokenGenerator {
	return &TokenGenerator{
		secretKey: cfg.SecretKey,
	}
}

func (g *TokenGenerator) GenerateNewAccessToken(params Params) (string, error) {
	claims := &Claims{
		Id: params.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(params.Ttl).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(g.secretKey))
}
