package mysql

import (
	"NihiStore/server/cmd/goods/config"
	"NihiStore/server/shared/kitex_gen/base"
	"NihiStore/server/shared/kitex_gen/goods"
	"NihiStore/server/shared/model"
	"fmt"
)

type MysqlGenerator struct{}

func (*MysqlGenerator) CreateGoods(goods *model.Goods) {
	config.DB.Create(goods)
}

func (*MysqlGenerator) SelectGoodsByUserIdAndGoodsId(SellerId, GoodsId int64) *model.Goods {
	var goods model.Goods
	fmt.Println("DEBUG")
	fmt.Println(SellerId, GoodsId)
	config.DB.Where("user_id = ? AND id = ?", SellerId, GoodsId).First(&goods)
	return &goods
}

func (*MysqlGenerator) DeleteGoods(goods *model.Goods) {
	config.DB.Unscoped().Delete(goods)
}

func (*MysqlGenerator) SearchGoodsInfo(req *goods.SearchGoodsInfoRequest) *[]model.Goods {
	var goods []model.Goods
	limit := req.PageAmount
	offset := (req.Page - 1) * limit
	config.DB.Offset(int(offset)).Limit(int(limit)).Where("name LIKE ? OR description LIKE ?", "%"+req.SearchMsg+"%", "%"+req.SearchMsg+"%").Find(&goods)
	return &goods
}

func (*MysqlGenerator) SelectGoodsById(Id int64) *model.Goods {
	var Goods model.Goods
	config.DB.Where("id = ?", Id).First(&Goods)
	return &Goods
}

func (*MysqlGenerator) UpdateGoods(ID int64, Goods *base.Goods) {
	config.DB.Model(&model.Goods{}).Where("id = ?", ID).Updates(Goods)
}
func (*MysqlGenerator) UpdateGoodsPhoto(goodsId, photoId int64) error {
	err := config.DB.Model(&model.Goods{}).Where("id = ?", goodsId).Update("photo_id", photoId).Error
	return err
}
func (*MysqlGenerator) GetPhotoId(goodsId int64) int64 {
	var goods model.Goods
	config.DB.Model(&model.Goods{}).Where("id = ?", goodsId).First(&goods)
	return goods.PhotoId
}
