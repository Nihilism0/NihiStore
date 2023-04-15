package main

import (
	pay "NihiStore/server/shared/kitex_gen/pay"
	"NihiStore/server/shared/tools"
	"context"
	"net/http"
	"net/url"
)

// PayServiceImpl implements the last service interface defined in the IDL.
type PayServiceImpl struct {
	MysqlPayGenerator
	ParseGenerator
}

type MysqlPayGenerator interface {
	GetSellerAliId(ID int64) string
}

type ParseGenerator interface {
	GetTradeNo(userId, GoodsId int64) string
	GetUrl(in *pay.BuyGoodsRequest, tradeNo, SellerAliId string) *url.URL
}

// BuyGoods implements the PayServiceImpl interface.
func (s *PayServiceImpl) BuyGoods(ctx context.Context, req *pay.BuyGoodsRequest) (resp *pay.BuyGoodsResponse, err error) {
	resp = new(pay.BuyGoodsResponse)
	var tradeNo = s.ParseGenerator.GetTradeNo(req.UserId, req.GoodsId)
	SellerAliId := s.MysqlPayGenerator.GetSellerAliId(req.UserId)
	URL := s.ParseGenerator.GetUrl(req, tradeNo, SellerAliId)
	urlStr := URL.String()
	resp.URL = urlStr
	resp.BaseResp = tools.BuildBaseResp(http.StatusOK, "Get pay URL success")
	return resp, nil
}
