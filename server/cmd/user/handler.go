package main

import (
	"NihiStore/server/cmd/user/config"
	"NihiStore/server/shared/consts"
	"NihiStore/server/shared/errx"
	user "NihiStore/server/shared/kitex_gen/user"
	"NihiStore/server/shared/model"
	"NihiStore/server/shared/tools"
	"context"
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
	resp = new(user.LoginResponse)
	theuser := model.User{}
	config.DB.Where("username = ?", req.Username).First(&theuser)
	if theuser.Username == "" {
		resp.BaseResp = tools.BuildBaseResp(errx.FindNone, "No such person found")
		return resp, nil
	}
	if theuser.Password != req.Password {
		resp.BaseResp = tools.BuildBaseResp(errx.PassWordWrong, "Wrong Password")
		return resp, nil
	}
	now := time.Now().Unix()
	tools.BuildToken()
	resp.Token, err = s.TokenGenerator.CreateToken(&model.CustomClaims{
		ID:       theuser.ID,
		IsSeller: theuser.IsSeller,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now,
			NotBefore: now,
			ExpiresAt: now + consts.ThirtyDays,
			Issuer:    consts.JWTIssuer,
		},
	},
	)
	resp.BaseResp = tools.BuildBaseResp(200, "Login Success")
	return resp, nil
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// TODO: Your code here...
	return
}
