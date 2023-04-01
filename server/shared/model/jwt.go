package model

import (
	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	ID       uint
	IsSeller bool
	jwt.StandardClaims
}
