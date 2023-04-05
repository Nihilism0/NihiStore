// Code generated by hertz generator.

package user

import (
	base "NihiStore/server/cmd/api/biz/model/base"
	huser "NihiStore/server/cmd/api/biz/model/user"
	"NihiStore/server/cmd/api/config"
	kuser "NihiStore/server/shared/kitex_gen/user"
	"context"
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
	var req user.CreateFavoritesReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// WatchFavorites .
// @router /user/favorites/watchfavorites [GET]
func WatchFavorites(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.WatchFavoritesReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// WatchGoodsInFavorites .
// @router /user/favorites/watchgoodsinfavorites [GET]
func WatchGoodsInFavorites(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.WatchGoodsInFavoritesReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// DeleteFavorites .
// @router /user/favorites/deletefavorites [DELETE]
func DeleteFavorites(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DeleteFavoritesReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// CollectGoods .
// @router /user/favorites/collectgoods [POST]
func CollectGoods(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.CollectGoodsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// AddToCart .
// @router /user/cart/addtocart [POST]
func AddToCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.AddToCartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// AddAmounrCart .
// @router /user/cart/addamountcart [POST]
func AddAmounrCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.AddAmountCartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// DeleteAmountCart .
// @router /user/cart/deleteamountcart [DELETE]
func DeleteAmountCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DeleteAmountCartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// RemoveOutCart .
// @router /user/cart/removeoutcart [DELETE]
func RemoveOutCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RemoveOutCartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// WatchCart .
// @router /user/cart/watchcart [GET]
func WatchCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.WatchCartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// CleanCart .
// @router /user/cart/cleancart [DELETE]
func CleanCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.CleanCartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}
