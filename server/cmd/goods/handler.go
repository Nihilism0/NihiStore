package main

import (
	"NihiStore/server/cmd/goods/config"
	"NihiStore/server/cmd/goods/dao"
	"NihiStore/server/shared/errx"
	goods "NihiStore/server/shared/kitex_gen/goods"
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
	// TODO: Your code here...
	return
}

// SearchGoodsInfo implements the GoodsServiceImpl interface.
func (s *GoodsServiceImpl) SearchGoodsInfo(ctx context.Context, req *goods.SearchGoodsInfoRequest) (resp *goods.SearchGoodsInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// SearchGoods implements the GoodsServiceImpl interface.
func (s *GoodsServiceImpl) SearchGoods(ctx context.Context, req *goods.SearchGoodsRequest) (resp *goods.SearchGoodsResponse, err error) {
	// TODO: Your code here...
	return
}
