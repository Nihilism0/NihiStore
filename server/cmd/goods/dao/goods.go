package dao

import (
	"NihiStore/server/shared/kitex_gen/goods"
	"NihiStore/server/shared/model"
)

func FillGoods(req *goods.CreateGoodsRequest) *model.Goods {
	return &model.Goods{
		Name:        req.GoodsInformation.Name,
		UserId:      req.Id,
		Describe:    req.GoodsInformation.Describe,
		Cost:        req.GoodsInformation.Cost,
		SalesVolume: 0,
	}
}
