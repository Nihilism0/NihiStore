package jwt

import (
	"NihiStore/server/cmd/user/config"
	"NihiStore/server/shared/consts"
	"NihiStore/server/shared/model"
	"github.com/golang-jwt/jwt"
	"time"
)

type TokenGenerator struct{}

func (TG *TokenGenerator) CreateJWTtoken(ID int64, isSeller bool) (string, error) {
	now := time.Now().Unix()
	claim := model.CustomClaims{
		ID:       ID,
		IsSeller: isSeller,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now,
			NotBefore: now,
			Issuer:    consts.JWTIssuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(config.GlobalServerConfig.JWTInfo.SigningKey))
}
