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
