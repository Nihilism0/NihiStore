package config

import (
	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

var (
	DB                 *gorm.DB
	GlobalServerConfig ServerConfig
	GlobalNacosConfig  NacosConfig
	Enforcer           *casbin.Enforcer
)
