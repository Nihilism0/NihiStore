package mysql

import (
	"NihiStore/server/cmd/pay/config"
	"NihiStore/server/shared/model"
)

type MysqlPayGenerator struct{}

func (*MysqlPayGenerator) GetSellerAliId(ID int64) string {
	var user model.User
	config.DB.Model(&user).Where("id = ?", ID).First(&user)
	return user.SellerAliId
}
