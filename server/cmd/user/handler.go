package main

import (
	"NihiStore/server/cmd/user/config"
	"NihiStore/server/shared/consts"
	"NihiStore/server/shared/errx"
	"NihiStore/server/shared/kitex_gen/base"
	user "NihiStore/server/shared/kitex_gen/user"
	"NihiStore/server/shared/model"
	"NihiStore/server/shared/tools"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
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
			Issuer:    consts.JWTIssuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	resp.Token, err = token.SignedString([]byte(config.GlobalServerConfig.JWTInfo.SigningKey))
	if err != nil {
		fmt.Println(err)
	}
	resp.BaseResp = tools.BuildBaseResp(200, "Login Success")
	return resp, nil
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
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
	resp = new(user.CreateFavoritesResponse)
	var Favorites model.Favorites
	config.DB.Where("id = ? AND name = ?", req.UserId, req.FavoritesName).First(&Favorites)
	if Favorites.Name != "" {
		resp.BaseResp = tools.BuildBaseResp(2, "Favorites already exist")
		return resp, nil
	}
	Favorites.Name = req.FavoritesName
	Favorites.Describe = req.Describe
	Favorites.UserId = req.UserId
	config.DB.Create(&Favorites)
	resp.BaseResp = tools.BuildBaseResp(200, "Create success")
	return resp, nil
}

// WatchFavorites implements the UserServiceImpl interface.
func (s *UserServiceImpl) WatchFavorites(ctx context.Context, req *user.WatchFavoritesRequset) (resp *user.WatchFavoritesResponse, err error) {
	resp = new(user.WatchFavoritesResponse)
	var favoriteses []model.Favorites
	config.DB.Where("id = ?", req.UserId).Find(&favoriteses)
	for _, v := range favoriteses {
		resp.Favoriteses = append(resp.Favoriteses, &base.Favorites{
			Name:     v.Name,
			Describe: v.Describe,
			Id:       int64(v.ID),
		})
	}
	resp.BaseResp = tools.BuildBaseResp(200, "Find success")
	return resp, nil
}

// WatchGoodsInFavorites implements the UserServiceImpl interface.
func (s *UserServiceImpl) WatchGoodsInFavorites(ctx context.Context, req *user.WatchGoodsInFavoritesRequest) (resp *user.WatchGoodsInFavoritesResponse, err error) {
	resp = new(user.WatchGoodsInFavoritesResponse)
	var collections []model.Collection
	config.DB.Where("user_id = ? AND favorites_id = ?", req.UserId, req.FavoritesId).Find(&collections)
	for _, v := range collections {
		resp.GoodsIds = append(resp.GoodsIds, &base.GoodsInFavorites{
			GoodsId: v.GoodsId,
		})
	}
	resp.BaseResp = tools.BuildBaseResp(200, "Find success")
	return resp, nil
}

// DeleteFavorites implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteFavorites(ctx context.Context, req *user.DeleteFavoritesRequest) (resp *user.DeleteFavoritesResponse, err error) {
	resp = new(user.DeleteFavoritesResponse)
	var favorites model.Favorites
	config.DB.Where("id = ? AND user_id = ?", req.FavoritesId, req.UserId).First(&favorites)
	if favorites.UserId != req.UserId {
		resp.BaseResp = tools.BuildBaseResp(errx.NoSuchFavorites, "No such favorites")
		return resp, nil
	}
	config.DB.Unscoped().Delete(&favorites)
	resp.BaseResp = tools.BuildBaseResp(200, "delete success")
	return resp, nil
}

// CollectGoods implements the UserServiceImpl interface.
func (s *UserServiceImpl) CollectGoods(ctx context.Context, req *user.CollectGoodsRequest) (resp *user.CollectGoodsResponse, err error) {
	resp = new(user.CollectGoodsResponse)
	var collection model.Collection
	var favorites model.Favorites
	config.DB.Where("id = ? AND user_id = ?", req.FavoritesId, req.UserId).First(&favorites)
	if favorites.UserId != req.UserId {
		resp.BaseResp = tools.BuildBaseResp(errx.FavoritesAuthFail, "Auth favorites user fail")
		return resp, nil
	}
	config.DB.Where("favorites_id = ? AND user_id = ? AND goods_id = ?", req.FavoritesId, req.UserId, req.GoodsId).First(&collection)
	if collection.GoodsId == req.GoodsId {
		resp.BaseResp = tools.BuildBaseResp(errx.AlreadyExist, "Goods already exist")
		return resp, nil
	}
	collection.GoodsId = req.GoodsId
	collection.FavoritesId = req.FavoritesId
	collection.UserId = req.UserId
	err = config.DB.Create(&collection).Error
	if err != nil {
		klog.Error(err)
		resp.BaseResp = tools.BuildBaseResp(500, "create fail")
		return resp, nil
	}
	resp.BaseResp = tools.BuildBaseResp(200, "Collect goods success")
	return resp, nil
}

// AddToCart implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddToCart(ctx context.Context, req *user.AddToCartRequest) (resp *user.AddToCartResponse, err error) {
	resp = new(user.AddToCartResponse)
	var cart model.Cart
	config.DB.Where("user_id = ? AND goods_id = ?", req.UserId, req.GoodsId).First(&cart)
	if cart.UserId == req.UserId {
		resp.BaseResp = tools.BuildBaseResp(errx.AlreadyExist, "Goods already in cart")
		return resp, nil
	}
	cart.GoodsId = req.GoodsId
	cart.UserId = req.UserId
	cart.Amount = req.Amount
	config.DB.Create(&cart)
	resp.BaseResp = tools.BuildBaseResp(200, "Add to cart success")
	return resp, nil
}

// AddAmountCart implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddAmountCart(ctx context.Context, req *user.AddAmountCartRequest) (resp *user.AddAmountCartResponse, err error) {
	resp = new(user.AddAmountCartResponse)
	var cart model.Cart
	config.DB.Where("user_id = ? AND goods_id = ?", req.UserId, req.GoodsId).First(&cart)
	if cart.UserId != req.UserId {
		resp.BaseResp = tools.BuildBaseResp(errx.NoSuchGoodsInCart, "No such goods in cart")
		return resp, nil
	}
	newamount := cart.Amount + req.Amount
	if (newamount) > consts.MaxGoodsAmount {
		resp.BaseResp = tools.BuildBaseResp(errx.OutOfMax, "Out of max")
		return resp, nil
	}
	config.DB.Model(&model.Cart{}).Where("user_id = ? AND goods_id = ?", req.UserId, req.GoodsId).Update("amount", newamount)
	resp.BaseResp = tools.BuildBaseResp(200, "Add cart amount success")
	return resp, nil
}

// DeleteAmountCart implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteAmountCart(ctx context.Context, req *user.DeleteAmountCartRequest) (resp *user.DeleteAmountCartResponse, err error) {
	resp = new(user.DeleteAmountCartResponse)
	var cart model.Cart
	config.DB.Where("user_id = ? AND goods_id = ?", req.UserId, req.GoodsId).First(&cart)
	if cart.UserId != req.UserId {
		resp.BaseResp = tools.BuildBaseResp(errx.NoSuchGoodsInCart, "No such goods in cart")
		return resp, nil
	}
	newamount := cart.Amount - req.Amount
	if newamount < 1 {
		resp.BaseResp = tools.BuildBaseResp(errx.OutOfMin, "Out of min")
		return resp, nil
	}
	config.DB.Model(&model.Cart{}).Where("user_id = ? AND goods_id = ?", req.UserId, req.GoodsId).Update("amount", newamount)
	resp.BaseResp = tools.BuildBaseResp(200, "Delete cart amount success")
	return resp, nil
}

// RemoveOutCart implements the UserServiceImpl interface.
func (s *UserServiceImpl) RemoveOutCart(ctx context.Context, req *user.RemoveOutCartRequest) (resp *user.RemoveOutCartResponse, err error) {
	resp = new(user.RemoveOutCartResponse)
	var cart model.Cart
	config.DB.Where("user_id = ? AND goods_id = ?", req.UserId, req.GoodsId).First(&cart)
	if cart.UserId != req.UserId {
		resp.BaseResp = tools.BuildBaseResp(errx.NoSuchGoodsInCart, "No such goods in cart")
		return resp, nil
	}
	config.DB.Unscoped().Delete(&cart)
	resp.BaseResp = tools.BuildBaseResp(200, "Delete cart success")
	return resp, nil
}

// WatchCart implements the UserServiceImpl interface.
func (s *UserServiceImpl) WatchCart(ctx context.Context, req *user.WatchCartRequest) (resp *user.WatchCartResponse, err error) {
	resp = new(user.WatchCartResponse)
	var carts []model.Cart
	config.DB.Where("user_id = ?", req.UserId).Find(&carts)
	resp.BaseResp = tools.BuildBaseResp(200, "Watch cart success")
	for _, v := range carts {
		resp.Goods = append(resp.Goods, &base.GoodsInCart{
			Amount:  v.Amount,
			GoodsId: v.GoodsId,
			UserId:  v.UserId,
		})
	}
	return resp, nil
}

// CleanCart implements the UserServiceImpl interface.
func (s *UserServiceImpl) CleanCart(ctx context.Context, req *user.CleanCartRequest) (resp *user.CleanCartResponse, err error) {
	resp = new(user.CleanCartResponse)
	config.DB.Unscoped().Where("user_id = ?", req.UserId).Delete(&model.Collection{})
	resp.BaseResp = tools.BuildBaseResp(200, "Clear success")
	return resp, nil
}
