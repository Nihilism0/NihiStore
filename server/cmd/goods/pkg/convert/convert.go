package convert

import (
	"NihiStore/server/shared/kitex_gen/base"
	"NihiStore/server/shared/kitex_gen/goods"
	"NihiStore/server/shared/model"
)

type ConvertGenerator struct{}

func (*ConvertGenerator) ConvertGoods(req *goods.CreateGoodsRequest) *model.Goods {
	return &model.Goods{
		Name:        req.GoodsInformation.Name,
		UserId:      req.Id,
		Description: req.GoodsInformation.Description,
		Cost:        req.GoodsInformation.Cost,
		SalesVolume: 0,
	}
}

func (*ConvertGenerator) ConvertGoodsFullInfo(Goods *model.Goods) *base.GoodsFullInfo {
	return &base.GoodsFullInfo{
		Id:          int64(Goods.ID),
		Name:        Goods.Name,
		UserId:      Goods.UserId,
		Description: Goods.Description,
		Cost:        Goods.Cost,
		SalesVolume: Goods.SalesVolume,
	}
}
