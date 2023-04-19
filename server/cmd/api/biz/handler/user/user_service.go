// Code generated by hertz generator.

package user

import (
	huser "NihiStore/server/cmd/api/biz/model/user"
	"NihiStore/server/cmd/api/config"
	kuser "NihiStore/server/shared/kitex_gen/user"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"net/http"
)

// Login .
// @router /login [GET]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req huser.LoginReq
	resp := new(kuser.LoginResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err = config.GlobalUserClient.Login(ctx, &kuser.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		hlog.Error("rpc user service err!", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// Register .
// @router /request [GET]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req huser.RegisterReq
	resp := new(kuser.RegisterResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err = config.GlobalUserClient.Register(ctx, &kuser.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		hlog.Error("rpc user service err!", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// CreateFavorites .
// @router /user/favorites/createfavorites [POST]
func CreateFavorites(ctx context.Context, c *app.RequestContext) {
	var err error
	var req huser.CreateFavoritesReq
	resp := new(kuser.CreateFavoritesResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	ID, _ := c.Get("ID")
	resp, err = config.GlobalUserClient.CreateFavorites(ctx, &kuser.CreateFavoritesRequest{
		FavoritesName: req.FavoritesName,
		Description:   req.Description,
		UserId:        ID.(int64),
	})
	if err != nil {
		hlog.Error("rpc user service err!", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// WatchFavorites .
// @router /user/favorites/watchfavorites [GET]
func WatchFavorites(ctx context.Context, c *app.RequestContext) {
	var err error
	var req huser.WatchFavoritesReq
	resp := new(kuser.WatchFavoritesResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	ID, _ := c.Get("ID")
	resp, err = config.GlobalUserClient.WatchFavorites(ctx, &kuser.WatchFavoritesRequset{
		UserId: ID.(int64),
	})
	if err != nil {
		hlog.Error("rpc user service err!", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// WatchGoodsInFavorites .
// @router /user/favorites/watchgoodsinfavorites [GET]
func WatchGoodsInFavorites(ctx context.Context, c *app.RequestContext) {
	var err error
	var req huser.WatchGoodsInFavoritesReq
	resp := new(kuser.WatchGoodsInFavoritesResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	ID, _ := c.Get("ID")
	resp, err = config.GlobalUserClient.WatchGoodsInFavorites(ctx, &kuser.WatchGoodsInFavoritesRequest{
		FavoritesId: req.FavoritesId,
		UserId:      ID.(int64),
	})
	if err != nil {
		hlog.Error("rpc user service err!", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// DeleteFavorites .
// @router /user/favorites/deletefavorites [DELETE]
func DeleteFavorites(ctx context.Context, c *app.RequestContext) {
	var err error
	var req huser.DeleteFavoritesReq
	resp := new(kuser.DeleteFavoritesResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	ID, _ := c.Get("ID")
	resp, err = config.GlobalUserClient.DeleteFavorites(ctx, &kuser.DeleteFavoritesRequest{
		FavoritesId: req.FavoritesId,
		UserId:      ID.(int64),
	})
	if err != nil {
		hlog.Error("rpc user service err!", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// CollectGoods .
// @router /user/favorites/collectgoods [POST]
func CollectGoods(ctx context.Context, c *app.RequestContext) {
	var err error
	var req huser.CollectGoodsReq
	resp := new(kuser.CollectGoodsResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	ID, _ := c.Get("ID")
	resp, err = config.GlobalUserClient.CollectGoods(ctx, &kuser.CollectGoodsRequest{
		GoodsId:     req.GoodsId,
		FavoritesId: req.FavoritesId,
		UserId:      ID.(int64),
	})
	fmt.Println("GET")
	fmt.Println(resp)
	if err != nil {
		hlog.Error("rpc user service err!", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// AddToCart .
// @router /user/cart/addtocart [POST]
func AddToCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req huser.AddToCartReq
	resp := new(kuser.AddToCartResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	ID, _ := c.Get("ID")
	resp, err = config.GlobalUserClient.AddToCart(ctx, &kuser.AddToCartRequest{
		GoodsId: req.GoodsId,
		Amount:  req.Amount,
		UserId:  ID.(int64),
	})
	if err != nil {
		hlog.Error("rpc user service err!", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// AddAmounrCart .
// @router /user/cart/addamountcart [POST]
func AddAmounrCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req huser.AddAmountCartReq
	resp := new(kuser.AddAmountCartResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	ID, _ := c.Get("ID")
	resp, err = config.GlobalUserClient.AddAmountCart(ctx, &kuser.AddAmountCartRequest{
		GoodsId: req.GoodsId,
		Amount:  req.Amount,
		UserId:  ID.(int64),
	})
	if err != nil {
		hlog.Error("rpc user service err!", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// DeleteAmountCart .
// @router /user/cart/deleteamountcart [DELETE]
func DeleteAmountCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req huser.DeleteAmountCartReq
	resp := new(kuser.DeleteAmountCartResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	ID, _ := c.Get("ID")
	resp, err = config.GlobalUserClient.DeleteAmountCart(ctx, &kuser.DeleteAmountCartRequest{
		GoodsId: req.GoodsId,
		Amount:  req.Amount,
		UserId:  ID.(int64),
	})
	if err != nil {
		hlog.Error("rpc user service err!", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// RemoveOutCart .
// @router /user/cart/removeoutcart [DELETE]
func RemoveOutCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req huser.RemoveOutCartReq
	resp := new(kuser.RemoveOutCartResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	ID, _ := c.Get("ID")
	resp, err = config.GlobalUserClient.RemoveOutCart(ctx, &kuser.RemoveOutCartRequest{
		GoodsId: req.GoodsId,
		UserId:  ID.(int64),
	})
	if err != nil {
		hlog.Error("rpc user service err!", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// WatchCart .
// @router /user/cart/watchcart [GET]
func WatchCart(ctx context.Context, c *app.RequestContext) {
	var err error
	resp := new(kuser.WatchCartResponse)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	ID, _ := c.Get("ID")
	resp, err = config.GlobalUserClient.WatchCart(ctx, &kuser.WatchCartRequest{
		UserId: ID.(int64),
	})
	if err != nil {
		hlog.Error("rpc user service err!", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// CleanCart .
// @router /user/cart/cleancart [DELETE]
func CleanCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req huser.CleanCartReq
	resp := new(kuser.CleanCartResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	ID, _ := c.Get("ID")
	resp, err = config.GlobalUserClient.CleanCart(ctx, &kuser.CleanCartRequest{UserId: ID.(int64)})
	if err != nil {
		hlog.Error("rpc user service err!", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// BeSeller .
// @router /user/beseller [PUT]
func BeSeller(ctx context.Context, c *app.RequestContext) {
	var err error
	var req huser.BeSellerReq
	resp := new(kuser.BeSellerResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	ID, _ := c.Get("ID")
	resp, err = config.GlobalUserClient.BeSeller(ctx, &kuser.BeSellerRequest{
		UserId:      ID.(int64),
		SellerAliId: req.SellerAliId,
	})
	if err != nil {
		hlog.Error("rpc user service err!", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetSellerByGoods .
// @router /user/getsellerbygoods [GET]
func GetSellerByGoods(ctx context.Context, c *app.RequestContext) {
	var err error
	var req huser.GetSellerByGoodsReq
	resp := new(kuser.GetSellerByGoodsResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err = config.GlobalUserClient.GetSellerByGoods(ctx, &kuser.GetSellerByGoodsRequest{
		GoodsId: req.GoodsId,
	})
	if err != nil {
		hlog.Error("rpc user service err!", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// UploadHead .
// @router /user/uploadhead [POST]
func UploadHead(ctx context.Context, c *app.RequestContext) {
	var err error
	resp := new(kuser.UploadHeadResponse)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	ID, _ := c.Get("ID")
	resp, err = config.GlobalUserClient.UploadHead(ctx, &kuser.UploadHeadRequest{
		UserId: ID.(int64),
	})
	if err != nil {
		hlog.Error("rpc user service err!", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// GetHead .
// @router /user/gethead [GET]
func GetHead(ctx context.Context, c *app.RequestContext) {
	var err error
	resp := new(kuser.GetHeadRespnse)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	ID, _ := c.Get("ID")
	resp, err = config.GlobalUserClient.GetHead(ctx, &kuser.GetHeadRequest{
		UserId: ID.(int64),
	})
	if err != nil {
		hlog.Error("rpc user service err!", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}
