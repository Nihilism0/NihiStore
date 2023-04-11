package mysql

import (
	"NihiStore/server/cmd/user/config"
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

func (*MysqlUserGenerator) BeSeller(UserId int64) {
	config.DB.Model(&model.User{}).Where("id = ?", UserId).Update("is_seller", true)
}
