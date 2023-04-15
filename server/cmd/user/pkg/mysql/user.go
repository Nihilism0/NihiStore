package mysql

import (
	"NihiStore/server/cmd/user/config"
	"NihiStore/server/shared/kitex_gen/user"
	"NihiStore/server/shared/model"
)

type MysqlUserGenerator struct{}

func (*MysqlUserGenerator) SelectUserFromUsername(username string) (*model.User, error) {
	theUser := model.User{}
	err := config.DB.Where("username = ?", username).First(&theUser).Error
	return &theUser, err
}

func (*MysqlUserGenerator) CreateUser(theUser *model.User) {
	config.DB.Create(theUser)
}

func (*MysqlUserGenerator) BeSellerBeSeller(in *user.BeSellerRequest) {
	config.DB.Model(&model.User{}).Where("id = ?", in.UserId).Updates(&model.User{IsSeller: true, SellerAliId: in.SellerAliId})
}
