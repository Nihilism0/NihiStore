package main

import (
	"NihiStore/server/cmd/goods/config"
	"NihiStore/server/cmd/goods/dao"
	"NihiStore/server/shared/errx"
	"NihiStore/server/shared/kitex_gen/base"
	goods "NihiStore/server/shared/kitex_gen/goods"
	"NihiStore/server/shared/model"
	"NihiStore/server/shared/tools"
	"context"
)

// GoodsServiceImpl implements the last service interface defined in the IDL.
type GoodsServiceImpl struct{}

// CreateGoods implements the GoodsServiceImpl interface.
func (s *GoodsServiceImpl) CreateGoods(ctx context.Context, req *goods.CreateGoodsRequest) (resp *goods.CreateGoodsResponse, err error) {
	resp = new(goods.CreateGoodsResponse)
	good := dao.FillGoods(req)
	err = config.DB.Create(&good).Error
	if err != nil {
		resp.BaseResp = tools.BuildBaseResp(errx.CreateGoodsFail, "Create Goods Fail")
		return resp, nil
	}
	resp.BaseResp = tools.BuildBaseResp(200, "Create goods success")
	return resp, nil
}

// DeleteGoods implements the GoodsServiceImpl interface.
func (s *GoodsServiceImpl) DeleteGoods(ctx context.Context, req *goods.DeleteGoodsRequest) (resp *goods.DeleteGoodsResponse, err error) {
	resp = new(goods.DeleteGoodsResponse)
	var goods model.Goods
	config.DB.Where("user_id = ? AND goods_id = ?", req.SellerId, req.GoodsId).First(&goods)
	if goods.UserId != req.SellerId {
		resp.BaseResp = tools.BuildBaseResp(errx.AuthGoodsFail, "No such goods")
		return resp, nil
	}
	config.DB.Unscoped().Delete(&goods)
	resp.BaseResp = tools.BuildBaseResp(200, "Delete goods success")
	return resp, nil
}

// SearchGoodsInfo implements the GoodsServiceImpl interface.
func (s *GoodsServiceImpl) SearchGoodsInfo(ctx context.Context, req *goods.SearchGoodsInfoRequest) (resp *goods.SearchGoodsInfoResponse, err error) {
	resp = new(goods.SearchGoodsInfoResponse)
	var goods []model.Goods
	limit := 10
	offset := (100 - 1) * limit
	config.DB.Offset(offset).Limit(limit).Where("MATCH(name, describe) AGAINST( ? IN BOOLEAN MODE)", req.SearchMsg).Find(&goods)
	for _, v := range goods {
		resp.Names = append(resp.Names, &base.Name{Name: v.Name})
	}
	resp.BaseResp = tools.BuildBaseResp(200, "Search goods info success")
	return resp, nil
}

// SearchGoods implements the GoodsServiceImpl interface.
func (s *GoodsServiceImpl) SearchGoods(ctx context.Context, req *goods.SearchGoodsRequest) (resp *goods.SearchGoodsResponse, err error) {
	var goods model.Goods
	config.DB.Where("id = ?", req.Id).Find(&goods)
	resp.GoodsFI = &base.GoodsFullInfo{
		Id:          int64(goods.ID),
		Name:        goods.Name,
		UserId:      goods.UserId,
		Describe:    goods.Describe,
		Cost:        goods.Cost,
		SalesVolume: goods.SalesVolume,
	}
	resp.BaseResp = tools.BuildBaseResp(200, "Search fullInfo success")
	return resp, nil
}
