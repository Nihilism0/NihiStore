package main

import (
	"NihiStore/server/cmd/user/config"
	"NihiStore/server/shared/consts"
	"NihiStore/server/shared/errx"
	user "NihiStore/server/shared/kitex_gen/user"
	"NihiStore/server/shared/model"
	"NihiStore/server/shared/tools"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)
	theUser := model.User{}
	config.DB.Where("username = ?", req.Username).First(&theUser)
	if theUser.Username == "" {
		resp.BaseResp = tools.BuildBaseResp(errx.FindNone, "No such person found")
		return resp, nil
	}
	if theUser.Password != req.Password {
		resp.BaseResp = tools.BuildBaseResp(errx.PassWordWrong, "Wrong Password")
		return resp, nil
	}
	now := time.Now().Unix()
	//jwt
	claim := model.CustomClaims{
		ID:       int64(theUser.ID),
		IsSeller: theUser.IsSeller,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now,
			NotBefore: now,
			ExpiresAt: now + consts.ThirtyDays,
			Issuer:    consts.JWTIssuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	fmt.Println("DEBUG")
	fmt.Println(config.GlobalServerConfig.JWTInfo.SigningKey)
	resp.Token, err = token.SignedString([]byte(config.GlobalServerConfig.JWTInfo.SigningKey))
	if err != nil {
		fmt.Println(err)
	}
	resp.BaseResp = tools.BuildBaseResp(200, "Login Success")
	return resp, nil
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(user.RegisterResponse)
	theUser := model.User{}
	config.DB.Where("username = ?", req.Username).First(&theUser)
	if theUser.Username != "" {
		resp.BaseResp = tools.BuildBaseResp(errx.AlreadyExist, "Username already exist")
		resp.OK = false
		return resp, nil
	}
	theUser.Username = req.Username
	theUser.IsSeller = false
	theUser.Password = req.Password
	errMsg := config.DB.Create(&theUser).Error
	if err != nil {
		resp.BaseResp = tools.BuildBaseResp(errx.CreatUserFail, errMsg.Error())
		resp.OK = false
		return resp, nil
	}
	resp.BaseResp = tools.BuildBaseResp(200, "Create user success")
	resp.OK = true
	return resp, nil
}

// CreateFavorites implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateFavorites(ctx context.Context, req *user.CreateFavoritesRequest) (resp *user.CreateFavoritesResponse, err error) {
	// TODO: Your code here...
	return
}

// WatchFavorites implements the UserServiceImpl interface.
func (s *UserServiceImpl) WatchFavorites(ctx context.Context, req *user.WatchFavoritesRequset) (resp *user.WatchFavoritesResponse, err error) {
	// TODO: Your code here...
	return
}

// WatchGoodsInFavorites implements the UserServiceImpl interface.
func (s *UserServiceImpl) WatchGoodsInFavorites(ctx context.Context, req *user.WatchGoodsInFavoritesRequest) (resp *user.WatchGoodsInFavoritesResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteFavorites implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteFavorites(ctx context.Context, req *user.DeleteFavoritesRequest) (resp *user.DeleteFavoritesResponse, err error) {
	// TODO: Your code here...
	return
}

// CollectGoods implements the UserServiceImpl interface.
func (s *UserServiceImpl) CollectGoods(ctx context.Context, req *user.CollectGoodsRequest) (resp *user.CollectGoodsResponse, err error) {
	// TODO: Your code here...
	return
}

// AddToCart implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddToCart(ctx context.Context, req *user.AddToCartRequest) (resp *user.AddToCartResponse, err error) {
	// TODO: Your code here...
	return
}

// AddAmountCart implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddAmountCart(ctx context.Context, req *user.AddAmountCartRequest) (resp *user.AddAmountCartResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteAmountCart implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteAmountCart(ctx context.Context, req *user.DeleteAmountCartRequest) (resp *user.DeleteAmountCartResponse, err error) {
	// TODO: Your code here...
	return
}

// RemoveOutCart implements the UserServiceImpl interface.
func (s *UserServiceImpl) RemoveOutCart(ctx context.Context, req *user.RemoveOutCartRequest) (resp *user.RemoveOutCartResponse, err error) {
	// TODO: Your code here...
	return
}

// WatchCart implements the UserServiceImpl interface.
func (s *UserServiceImpl) WatchCart(ctx context.Context, req *user.AddToCartRequest) (resp *user.WatchCartResponse, err error) {
	// TODO: Your code here...
	return
}

// CleanCart implements the UserServiceImpl interface.
func (s *UserServiceImpl) CleanCart(ctx context.Context, req *user.CleanCartRequest) (resp *user.CleanCartResponse, err error) {
	// TODO: Your code here...
	return
}
