package main

import (
	"NihiStore/server/shared/errx"
	"NihiStore/server/shared/kitex_gen/base"
	goods "NihiStore/server/shared/kitex_gen/goods"
	"NihiStore/server/shared/model"
	"NihiStore/server/shared/tools"
	"context"
)

// GoodsServiceImpl implements the last service interface defined in the IDL.
type GoodsServiceImpl struct {
	ConvertGenerator
	MysqlGenerator
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
}

// CreateGoods implements the GoodsServiceImpl interface.
func (s *GoodsServiceImpl) CreateGoods(ctx context.Context, req *goods.CreateGoodsRequest) (resp *goods.CreateGoodsResponse, err error) {
	resp = new(goods.CreateGoodsResponse)
	Goods := s.ConvertGenerator.ConvertGoods(req)
	s.MysqlGenerator.CreateGoods(Goods)
	resp.BaseResp = tools.BuildBaseResp(200, "Create goods success")
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
	resp.BaseResp = tools.BuildBaseResp(200, "Delete goods success")
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
	resp.BaseResp = tools.BuildBaseResp(200, "Search full info success")
	return resp, nil
}

// UpdateGoods implements the GoodsServiceImpl interface.
func (s *GoodsServiceImpl) UpdateGoods(ctx context.Context, req *goods.UpdateGoodsRequest) (resp *goods.UpdateGoodsResponse, err error) {
	resp = new(goods.UpdateGoodsResponse)
	Goods := s.MysqlGenerator.SelectGoodsById(req.Id)
	if Goods.UserId != req.UserId {
		resp.BaseResp = tools.BuildBaseResp(errx.AuthSellerFail, "Goods is not belong to you")
		return resp, nil
	}
	s.MysqlGenerator.UpdateGoods(req.Id, req.GoodsInformation)
	resp.BaseResp = tools.BuildBaseResp(200, "Update success")
	return resp, nil
}
