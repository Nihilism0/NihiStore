package mysql

import (
	"NihiStore/server/cmd/user/config"
	"NihiStore/server/shared/kitex_gen/user"
	"NihiStore/server/shared/model"
	"NihiStore/server/shared/tools"
	"strconv"
)

type MysqlUserGenerator struct{}

func (*MysqlUserGenerator) SelectUserFromUsername(username string) (*model.User, error) {
	theUser := model.User{}
	err := config.DB.Where("username = ?", username).First(&theUser).Error
	return &theUser, err
}

func (*MysqlUserGenerator) CreateUser(theUser *model.User) {
	salt := tools.GenValidateCode(6)
	theUser.Password = tools.MakePassword(theUser.Password, salt)
	theUser.Salt = salt
	config.DB.Create(theUser)
}

func (*MysqlUserGenerator) BeSeller(in *user.BeSellerRequest) {
	config.DB.Model(&model.User{}).Where("id = ?", in.UserId).Updates(&model.User{SellerAliId: in.SellerAliId})
	config.Enforcer.AddRoleForUser(strconv.FormatInt(in.UserId, 10), "seller", "g")
}

func (*MysqlUserGenerator) GetSellerByGoods(goodsId int64) int64 {
	var goods model.Goods
	config.DB.Where("id = ?", goodsId).First(&goods)
	return goods.UserId
}

func (*MysqlUserGenerator) UpdateHeadId(userId, headId int64) error {
	err := config.DB.Model(&model.User{}).Where("id = ?", userId).Update("head_id", headId).Error
	return err
}

func (*MysqlUserGenerator) GetHeadId(userId int64) int64 {
	var user model.User
	config.DB.Model(&model.User{}).Where("id = ?", userId).First(&user)
	return user.HeadId
}
