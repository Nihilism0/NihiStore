package mysql

import (
	"NihiStore/server/cmd/user/config"
	"NihiStore/server/shared/model"
)

type MysqlCartGenerator struct{}

func (*MysqlCartGenerator) SelectCartByUserIdAndGoodsId(UserId, GoodsId int64) *model.Cart {
	var cart model.Cart
	config.DB.Where("user_id = ? AND goods_id = ?", UserId, GoodsId).First(&cart)
	return &cart
}

func (*MysqlCartGenerator) CreateCart(cart *model.Cart) {
	config.DB.Create(cart)
}

func (*MysqlCartGenerator) UpdateCart(UserId, GoodsId, newamount int64) {
	config.DB.Model(&model.Cart{}).Where("user_id = ? AND goods_id = ?", UserId, GoodsId).Update("amount", newamount)
}

func (*MysqlCartGenerator) RemoveCart(cart *model.Cart) {
	config.DB.Unscoped().Delete(cart)
}

func (*MysqlCartGenerator) SelectCartByUserId(UserId int64) *[]model.Cart {
	var carts []model.Cart
	config.DB.Where("user_id = ?", UserId).Find(&carts)
	return &carts
}

func (*MysqlCartGenerator) RemoveAllCart(UserId int64) {
	config.DB.Unscoped().Where("user_id = ?", UserId).Delete(&model.Cart{})
}
