package main

import (
	"NihiStore/server/cmd/goods/config"
	"NihiStore/server/shared/errx"
	"NihiStore/server/shared/kitex_gen/base"
	goods "NihiStore/server/shared/kitex_gen/goods"
	"NihiStore/server/shared/kitex_gen/oss"
	"NihiStore/server/shared/model"
	"NihiStore/server/shared/tools"
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	"net/http"
	"strconv"
	"time"
)

// GoodsServiceImpl implements the last service interface defined in the IDL.
type GoodsServiceImpl struct {
	ConvertGenerator
	MysqlGenerator
	OSSManager
}

type ConvertGenerator interface {
	ConvertGoods(req *goods.CreateGoodsRequest) *model.Goods
	ConvertGoodsFullInfo(Goods *model.Goods) *base.GoodsFullInfo
}

type MysqlGenerator interface {
	CreateGoods(goods *model.Goods)
	SelectGoodsByUserIdAndGoodsId(SellerId, GoodsId int64) *model.Goods
	DeleteGoods(goods *model.Goods)
	SearchGoodsInfo(req *goods.SearchGoodsInfoRequest) *[]model.Goods
	SelectGoodsById(Id int64) *model.Goods
	UpdateGoods(ID int64, Goods *base.Goods)
	UpdateGoodsPhoto(goodsId, photoId int64) error
	GetPhotoId(goodsId int64) int64
}

type OSSManager interface {
	CreateGoodsOSS(ctx context.Context, req *oss.CreateGoodsOSSRequest, callOptions ...callopt.Option) (r *oss.CreateGoodsOSSResponse, err error)
	GetGoodsOSS(ctx context.Context, req *oss.GetGoodsOSSRequest, callOptions ...callopt.Option) (r *oss.GetGoodsOSSResponse, err error)
}

// CreateGoods implements the GoodsServiceImpl interface.
func (s *GoodsServiceImpl) CreateGoods(ctx context.Context, req *goods.CreateGoodsRequest) (resp *goods.CreateGoodsResponse, err error) {
	resp = new(goods.CreateGoodsResponse)
	ok, err := config.Enforcer.Enforce(strconv.FormatInt(req.Id, 10), "create", "create")
	if err != nil {
		klog.Error(err)
		resp.BaseResp = tools.BuildBaseResp(errx.CasbinInternalError, "Enforce internal error")
		return resp, nil
	}
	if !ok {
		klog.Info(req.Id, " has no rules")
		resp.BaseResp = tools.BuildBaseResp(errx.CasbinAuthFail, "Casbin auth fail")
		return resp, nil
	}

	Goods := s.ConvertGenerator.ConvertGoods(req)
	s.MysqlGenerator.CreateGoods(Goods)
	config.Enforcer.AddPolicy(strconv.FormatInt(req.Id, 10), strconv.FormatInt(int64(Goods.ID), 10), "operate")
	resp.BaseResp = tools.BuildBaseResp(http.StatusOK, "Create goods success")
	return resp, nil
}

// DeleteGoods implements the GoodsServiceImpl interface.
func (s *GoodsServiceImpl) DeleteGoods(ctx context.Context, req *goods.DeleteGoodsRequest) (resp *goods.DeleteGoodsResponse, err error) {
	resp = new(goods.DeleteGoodsResponse)
	Goods := s.MysqlGenerator.SelectGoodsByUserIdAndGoodsId(req.SellerId, req.GoodsId)
	if Goods.UserId != req.SellerId {
		resp.BaseResp = tools.BuildBaseResp(errx.AuthGoodsFail, "No such goods")
		return resp, nil
	}
	s.MysqlGenerator.DeleteGoods(Goods)
	resp.BaseResp = tools.BuildBaseResp(http.StatusOK, "Delete goods success")
	return resp, nil
}

// SearchGoodsInfo implements the GoodsServiceImpl interface.
func (s *GoodsServiceImpl) SearchGoodsInfo(ctx context.Context, req *goods.SearchGoodsInfoRequest) (resp *goods.SearchGoodsInfoResponse, err error) {
	resp = new(goods.SearchGoodsInfoResponse)
	Goods := s.MysqlGenerator.SearchGoodsInfo(req)
	for _, v := range *Goods {
		resp.Names = append(resp.Names, &base.Name{
			Name: v.Name,
			Id:   int64(v.ID),
		})
	}
	resp.BaseResp = tools.BuildBaseResp(200, "Search goods info success")
	return resp, nil
}

// SearchGoods implements the GoodsServiceImpl interface.
func (s *GoodsServiceImpl) SearchGoods(ctx context.Context, req *goods.SearchGoodsRequest) (resp *goods.SearchGoodsResponse, err error) {
	resp = new(goods.SearchGoodsResponse)
	Goods := s.MysqlGenerator.SelectGoodsById(req.Id)
	resp.GoodsFI = s.ConvertGenerator.ConvertGoodsFullInfo(Goods)
	resp.BaseResp = tools.BuildBaseResp(http.StatusOK, "Search full info success")
	return resp, nil
}

// UpdateGoods implements the GoodsServiceImpl interface.
func (s *GoodsServiceImpl) UpdateGoods(ctx context.Context, req *goods.UpdateGoodsRequest) (resp *goods.UpdateGoodsResponse, err error) {
	resp = new(goods.UpdateGoodsResponse)
	ok, err := config.Enforcer.Enforce(strconv.FormatInt(req.UserId, 10), strconv.FormatInt(req.Id, 10), "operate")
	if err != nil {
		klog.Error(err)
		resp.BaseResp = tools.BuildBaseResp(errx.CasbinInternalError, "Enforce internal error")
		return resp, nil
	}
	if !ok {
		klog.Info(req.Id, " has no rules")
		resp.BaseResp = tools.BuildBaseResp(errx.CasbinAuthFail, "Casbin auth fail")
		return resp, nil
	}
	Goods := s.MysqlGenerator.SelectGoodsById(req.Id)
	if Goods.UserId != req.UserId {
		resp.BaseResp = tools.BuildBaseResp(errx.AuthSellerFail, "Goods is not belong to you")
		return resp, nil
	}
	s.MysqlGenerator.UpdateGoods(req.Id, req.GoodsInformation)
	resp.BaseResp = tools.BuildBaseResp(http.StatusOK, "Update success")
	return resp, nil
}

// UploadGoodsPhoto implements the GoodsServiceImpl interface.
func (s *GoodsServiceImpl) UploadGoodsPhoto(ctx context.Context, req *goods.UploadGoodsPhotoRequest) (resp *goods.UploadGoodsPhotoResponse, err error) {
	resp = new(goods.UploadGoodsPhotoResponse)
	path := tools.CreateGoodsPhotoMinioPath(strconv.FormatInt(req.GoodsId, 10))
	hr, err := s.OSSManager.CreateGoodsOSS(ctx, &oss.CreateGoodsOSSRequest{
		Path:       path,
		TimeoutSec: int32(10 * time.Second.Seconds()),
		GoodsId:    req.GoodsId,
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
	err = s.MysqlGenerator.UpdateGoodsPhoto(req.GoodsId, hr.Id)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = tools.BuildBaseResp(errx.UpdateGoodsSqlErr, err.Error())
		return resp, nil
	}
	resp.Url = hr.UploadUrl
	resp.BaseResp = tools.BuildBaseResp(http.StatusOK, "Upload photo url get success")
	return resp, nil
}

// GetGoodsPhoto implements the GoodsServiceImpl interface.
func (s *GoodsServiceImpl) GetGoodsPhoto(ctx context.Context, req *goods.GetGoodsPhotoRequest) (resp *goods.GetGoodsPhotoResponse, err error) {
	resp = new(goods.GetGoodsPhotoResponse)
	photoId := s.MysqlGenerator.GetPhotoId(req.GoodsId)
	hr, err := s.OSSManager.GetGoodsOSS(ctx, &oss.GetGoodsOSSRequest{
		Id:         photoId,
		TimeoutSec: int32(5 * time.Second.Seconds()),
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
	resp.BaseResp = tools.BuildBaseResp(http.StatusOK, "Get goods photo url success")
	return resp, nil
}
