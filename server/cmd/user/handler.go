package main

import (
	"NihiStore/server/cmd/user/config"
	"NihiStore/server/shared/consts"
	"NihiStore/server/shared/errx"
	"NihiStore/server/shared/kitex_gen/base"
	"NihiStore/server/shared/kitex_gen/oss"
	user "NihiStore/server/shared/kitex_gen/user"
	"NihiStore/server/shared/model"
	"NihiStore/server/shared/tools"
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	"net/http"
	"strconv"
	"time"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	TokenGenerator
	ConvertGenerator
	MysqlUserGenerator
	MysqlFavoGenerator
	MysqlCartGenerator
	OSSManager
}
type TokenGenerator interface {
	CreateJWTtoken(ID int64, isSeller bool) (string, error)
}

type ConvertGenerator interface {
	ConvertUser(req *user.RegisterRequest, user *model.User)
	ConvertFavorites(req *user.CreateFavoritesRequest, favorites *model.Favorites)
}

type MysqlUserGenerator interface {
	SelectUserFromUsername(username string) (*model.User, error)
	CreateUser(theUser *model.User)
	BeSeller(in *user.BeSellerRequest)
	GetSellerByGoods(goodsId int64) int64
	UpdateHeadId(userId, headId int64) error
	GetHeadId(userId int64) int64
}

type MysqlFavoGenerator interface {
	SelectFavoFromUserIdAndFavoName(UserId int64, FavoritesName string) *model.Favorites
	SelectFavoInUserId(UserID int64) *[]model.Favorites
	SelectCollectionsByUserAndFavo(UserId, FavoritesId int64) *[]model.Collection
	SelectFavoByIdAndUserId(FavoritesId, UserId int64) *model.Favorites
	SelectCollectionByAllId(FavoritesId, UserId, GoodsId int64) *model.Collection
	CreateCollection(collection *model.Collection)
}

type MysqlCartGenerator interface {
	SelectCartByUserIdAndGoodsId(UserId, GoodsId int64) *model.Cart
	CreateCart(cart *model.Cart)
	UpdateCart(UserId, GoodsId, newamount int64)
	RemoveCart(cart *model.Cart)
	SelectCartByUserId(UserId int64) *[]model.Cart
	RemoveAllCart(UserId int64)
}

type OSSManager interface {
	CreateHeadOSS(ctx context.Context, req *oss.CreateHeadOSSRequest, callOptions ...callopt.Option) (r *oss.CreateHeadOSSResponse, err error)
	GetHeadOSS(ctx context.Context, req *oss.GetHeadOSSRequest, callOptions ...callopt.Option) (r *oss.GetHeadOSSResponse, err error)
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)
	theUser, err := s.MysqlUserGenerator.SelectUserFromUsername(req.Username)
	if theUser.Username != req.Username {
		resp.BaseResp = tools.BuildBaseResp(errx.FindNone, "No such person found")
		return resp, nil
	}
	if theUser.Password != req.Password {
		resp.BaseResp = tools.BuildBaseResp(errx.PassWordWrong, "Wrong Password")
		return resp, nil
	}
	resp.Token, err = s.TokenGenerator.CreateJWTtoken(int64(theUser.ID), theUser.IsSeller)
	if err != nil {
		klog.Error("Create JWT error!", err)
		return
	}
	resp.BaseResp = tools.BuildBaseResp(200, "Login Success")
	return resp, nil
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp = new(user.RegisterResponse)
	theUser, err := s.MysqlUserGenerator.SelectUserFromUsername(req.Username)
	if theUser.Username == req.Username {
		resp.BaseResp = tools.BuildBaseResp(errx.AlreadyExist, "Username already exist")
		return resp, nil
	}
	s.ConvertGenerator.ConvertUser(req, theUser)
	s.MysqlUserGenerator.CreateUser(theUser)
	resp.BaseResp = tools.BuildBaseResp(200, "Create user success")
	return resp, nil
}

// CreateFavorites implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateFavorites(ctx context.Context, req *user.CreateFavoritesRequest) (resp *user.CreateFavoritesResponse, err error) {
	resp = new(user.CreateFavoritesResponse)
	favorites := s.MysqlFavoGenerator.SelectFavoFromUserIdAndFavoName(req.UserId, req.FavoritesName)
	if favorites.Name != "" {
		resp.BaseResp = tools.BuildBaseResp(2, "Favorites already exist")
		return resp, nil
	}

	favorites.Name = req.FavoritesName
	favorites.Description = req.Description
	favorites.UserId = req.UserId
	err = config.DB.Create(&favorites).Error
	if err != nil {

		return
	}
	resp.BaseResp = tools.BuildBaseResp(200, "Create success")
	return resp, nil
}

// WatchFavorites implements the UserServiceImpl interface.
func (s *UserServiceImpl) WatchFavorites(ctx context.Context, req *user.WatchFavoritesRequset) (resp *user.WatchFavoritesResponse, err error) {
	resp = new(user.WatchFavoritesResponse)
	favoriteses := s.MysqlFavoGenerator.SelectFavoInUserId(req.UserId)
	for _, v := range *favoriteses {
		resp.Favoriteses = append(resp.Favoriteses, &base.Favorites{
			Name:        v.Name,
			Description: v.Description,
			Id:          int64(v.ID),
		})
	}
	resp.BaseResp = tools.BuildBaseResp(200, "Find success")
	return resp, nil
}

// WatchGoodsInFavorites implements the UserServiceImpl interface.
func (s *UserServiceImpl) WatchGoodsInFavorites(ctx context.Context, req *user.WatchGoodsInFavoritesRequest) (resp *user.WatchGoodsInFavoritesResponse, err error) {
	resp = new(user.WatchGoodsInFavoritesResponse)
	collections := s.MysqlFavoGenerator.SelectCollectionsByUserAndFavo(req.UserId, req.FavoritesId)
	for _, v := range *collections {
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
	favorites := s.MysqlFavoGenerator.SelectFavoByIdAndUserId(req.FavoritesId, req.UserId)
	if favorites.UserId != req.UserId {
		resp.BaseResp = tools.BuildBaseResp(errx.NoSuchFavorites, "No such favorites")
		return resp, nil
	}

	err = config.DB.Unscoped().Delete(&favorites).Error
	if err != nil {
		klog.Error("Delete favorites err,", err)
		return
	}
	resp.BaseResp = tools.BuildBaseResp(200, "delete success")
	return resp, nil
}

// CollectGoods implements the UserServiceImpl interface.
func (s *UserServiceImpl) CollectGoods(ctx context.Context, req *user.CollectGoodsRequest) (resp *user.CollectGoodsResponse, err error) {
	resp = new(user.CollectGoodsResponse)
	favorites := s.MysqlFavoGenerator.SelectFavoByIdAndUserId(req.FavoritesId, req.UserId)
	if favorites.UserId != req.UserId {
		resp.BaseResp = tools.BuildBaseResp(errx.FavoritesAuthFail, "Auth favorites user fail")
		return resp, nil
	}
	collection := s.MysqlFavoGenerator.SelectCollectionByAllId(req.FavoritesId, req.UserId, req.GoodsId)
	if collection.GoodsId == req.GoodsId {
		resp.BaseResp = tools.BuildBaseResp(errx.AlreadyExist, "Goods already exist")
		return resp, nil
	}
	collection.GoodsId = req.GoodsId
	collection.FavoritesId = req.FavoritesId
	collection.UserId = req.UserId
	s.MysqlFavoGenerator.CreateCollection(collection)
	resp.BaseResp = tools.BuildBaseResp(200, "Collect goods success")
	return resp, nil
}

// AddToCart implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddToCart(ctx context.Context, req *user.AddToCartRequest) (resp *user.AddToCartResponse, err error) {
	resp = new(user.AddToCartResponse)
	cart := s.MysqlCartGenerator.SelectCartByUserIdAndGoodsId(req.UserId, req.GoodsId)
	if cart.UserId == req.UserId {
		resp.BaseResp = tools.BuildBaseResp(errx.AlreadyExist, "Goods already in cart")
		return resp, nil
	}
	cart.GoodsId = req.GoodsId
	cart.UserId = req.UserId
	cart.Amount = req.Amount
	s.MysqlCartGenerator.CreateCart(cart)
	resp.BaseResp = tools.BuildBaseResp(200, "Add to cart success")
	return resp, nil
}

// AddAmountCart implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddAmountCart(ctx context.Context, req *user.AddAmountCartRequest) (resp *user.AddAmountCartResponse, err error) {
	resp = new(user.AddAmountCartResponse)
	cart := s.MysqlCartGenerator.SelectCartByUserIdAndGoodsId(req.UserId, req.GoodsId)
	if cart.UserId != req.UserId {
		resp.BaseResp = tools.BuildBaseResp(errx.NoSuchGoodsInCart, "No such goods in cart")
		return resp, nil
	}
	newamount := cart.Amount + req.Amount
	if (newamount) > consts.MaxGoodsAmount {
		resp.BaseResp = tools.BuildBaseResp(errx.OutOfMax, "Out of max")
		return resp, nil
	}
	s.MysqlCartGenerator.UpdateCart(req.UserId, req.GoodsId, newamount)
	resp.BaseResp = tools.BuildBaseResp(200, "Add cart amount success")
	return resp, nil
}

// DeleteAmountCart implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteAmountCart(ctx context.Context, req *user.DeleteAmountCartRequest) (resp *user.DeleteAmountCartResponse, err error) {
	resp = new(user.DeleteAmountCartResponse)
	cart := s.MysqlCartGenerator.SelectCartByUserIdAndGoodsId(req.UserId, req.GoodsId)
	if cart.UserId != req.UserId {
		resp.BaseResp = tools.BuildBaseResp(errx.NoSuchGoodsInCart, "No such goods in cart")
		return resp, nil
	}
	newamount := cart.Amount - req.Amount
	if newamount < 1 {
		resp.BaseResp = tools.BuildBaseResp(errx.OutOfMin, "Out of min")
		return resp, nil
	}
	s.MysqlCartGenerator.UpdateCart(req.UserId, req.GoodsId, newamount)
	resp.BaseResp = tools.BuildBaseResp(200, "Delete cart amount success")
	return resp, nil
}

// RemoveOutCart implements the UserServiceImpl interface.
func (s *UserServiceImpl) RemoveOutCart(ctx context.Context, req *user.RemoveOutCartRequest) (resp *user.RemoveOutCartResponse, err error) {
	resp = new(user.RemoveOutCartResponse)
	cart := s.MysqlCartGenerator.SelectCartByUserIdAndGoodsId(req.UserId, req.GoodsId)
	if cart.UserId != req.UserId {
		resp.BaseResp = tools.BuildBaseResp(errx.NoSuchGoodsInCart, "No such goods in cart")
		return resp, nil
	}
	s.MysqlCartGenerator.RemoveCart(cart)
	resp.BaseResp = tools.BuildBaseResp(200, "Delete cart success")
	return resp, nil
}

// WatchCart implements the UserServiceImpl interface.
func (s *UserServiceImpl) WatchCart(ctx context.Context, req *user.WatchCartRequest) (resp *user.WatchCartResponse, err error) {
	resp = new(user.WatchCartResponse)
	carts := s.MysqlCartGenerator.SelectCartByUserId(req.UserId)
	resp.BaseResp = tools.BuildBaseResp(200, "Watch cart success")
	for _, v := range *carts {
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
	s.MysqlCartGenerator.RemoveAllCart(req.UserId)
	resp.BaseResp = tools.BuildBaseResp(http.StatusOK, "Clear success")
	return resp, nil
}

// BeSeller implements the UserServiceImpl interface.
func (s *UserServiceImpl) BeSeller(ctx context.Context, req *user.BeSellerRequest) (resp *user.BeSellerResponse, err error) {
	resp = new(user.BeSellerResponse)
	s.MysqlUserGenerator.BeSeller(req)
	resp.BaseResp = tools.BuildBaseResp(http.StatusOK, "Update seller success")
	return resp, nil
}

// GetSellerByGoods implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetSellerByGoods(ctx context.Context, req *user.GetSellerByGoodsRequest) (resp *user.GetSellerByGoodsResponse, err error) {
	resp = new(user.GetSellerByGoodsResponse)
	num, _ := strconv.ParseInt(req.GoodsId, 10, 64)
	resp.SellerId = strconv.FormatInt(s.MysqlUserGenerator.GetSellerByGoods(num), 10)
	resp.BaseResp.StatusCode = http.StatusOK
	resp.BaseResp.StatusMsg = "Find seller by goodsId success"
	return resp, nil
}

// UploadHead implements the UserServiceImpl interface.
func (s *UserServiceImpl) UploadHead(ctx context.Context, req *user.UploadHeadRequest) (resp *user.UploadHeadResponse, err error) {
	resp = new(user.UploadHeadResponse)
	path := tools.CreateHeadMinioPath(strconv.FormatInt(req.UserId, 10))
	hr, err := s.OSSManager.CreateHeadOSS(ctx, &oss.CreateHeadOSSRequest{
		Path:       path,
		TimeoutSec: int32(100 * time.Second.Seconds()),
		UserId:     req.UserId,
	})
	if err != nil {
		klog.Error(err)
		resp.BaseResp = tools.BuildBaseResp(http.StatusInternalServerError, err.Error())
		return resp, nil
	}
	if hr.BaseResp.StatusCode != http.StatusOK {
		resp.BaseResp = hr.BaseResp
		return resp, nil
	}
	err = s.MysqlUserGenerator.UpdateHeadId(req.UserId, hr.Id)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = tools.BuildBaseResp(errx.UpdateHeadIdErr, err.Error())
		return resp, nil
	}
	resp.Url = hr.UploadUrl
	resp.BaseResp = tools.BuildBaseResp(http.StatusOK, "Upload head url get success")
	return resp, nil
}

// GetHead implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetHead(ctx context.Context, req *user.GetHeadRequest) (resp *user.GetHeadRespnse, err error) {
	resp = new(user.GetHeadRespnse)
	headId := s.MysqlUserGenerator.GetHeadId(req.UserId)
	hr, err := s.OSSManager.GetHeadOSS(ctx, &oss.GetHeadOSSRequest{
		Id:         headId,
		TimeoutSec: int32(50 * time.Second.Seconds()),
	})
	if err != nil {
		klog.Error(err)
		resp.BaseResp = tools.BuildBaseResp(http.StatusInternalServerError, err.Error())
		return resp, nil
	}
	if hr.BaseResp.StatusCode != http.StatusOK {
		klog.Error(hr.BaseResp.StatusMsg)
		resp.BaseResp = hr.BaseResp
		return resp, nil
	}
	resp.Url = hr.Url
	resp.BaseResp = tools.BuildBaseResp(http.StatusOK, "Get head url success")
	return resp, nil
}
