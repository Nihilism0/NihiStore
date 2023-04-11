package config

import (
	"github.com/smartwalle/alipay/v3"
	"gorm.io/gorm"
)

var (
	DB                 *gorm.DB
	GlobalServerConfig ServerConfig
	GlobalNacosConfig  NacosConfig
	AliClient          *alipay.Client
)
