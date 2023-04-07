package mysql

import (
	"NihiStore/server/cmd/user/config"
	"NihiStore/server/shared/model"
)

type MysqlUserGenerator struct{}

func (*MysqlUserGenerator) SelectUserFromUsername(username string) model.User {
	theUser := model.User{}
	config.DB.Where("username = ?", username).First(&theUser)
	return theUser
}
