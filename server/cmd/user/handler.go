package main

import (
	"NihiStore/server/cmd/user/config"
	"NihiStore/server/shared/consts"
	"NihiStore/server/shared/errx"
	user "NihiStore/server/shared/kitex_gen/user"
	"NihiStore/server/shared/model"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	TokenGenerator
}

type TokenGenerator interface {
	CreateToken(claims *model.CustomClaims) (token string, err error)
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	fmt.Println("HHEHEHEHEHHEHE")
	config.DB.AutoMigrate(&model.User{})
	resp = new(user.LoginResponse)
	var user model.User
	config.DB.Where("username = ? ", req.Username).First(&user)
	if user.Username == "" {
		resp.BaseResp.StatusCode = errx.FindNone
		resp.BaseResp.StatusMsg = "No such person found"
		return
	}
	if user.Password != req.Password {
		resp.BaseResp.StatusCode = errx.PassWordWrong
		resp.BaseResp.StatusMsg = "Wrong Password"
		return
	}
	now := time.Now().Unix()
	resp.Token, err = s.TokenGenerator.CreateToken(&model.CustomClaims{
		ID:       user.ID,
		IsSeller: user.IsSeller,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now,
			NotBefore: now,
			ExpiresAt: now + consts.ThirtyDays,
			Issuer:    consts.JWTIssuer,
		},
	})
	resp.BaseResp.StatusCode = 200
	resp.BaseResp.StatusMsg = "Login Success"
	return
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// TODO: Your code here...
	return
}
