package jwt

import (
	"NihiStore/server/shared/middleware"
	"NihiStore/server/shared/model"
	"github.com/golang-jwt/jwt"
)

type TokenGenerator struct {
	jwt *middleware.JWT
}

func NewTokenGenerator(signingKey string) *TokenGenerator {
	j := middleware.NewJWT(signingKey)
	return &TokenGenerator{jwt: j}
}

func (g *TokenGenerator) CreateToken(claim *model.CustomClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(g.jwt.SigningKey)
	return tokenString, err
}
