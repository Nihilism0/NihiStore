package config

import "gorm.io/gorm"

var (
	DB                 *gorm.DB
	GlobalServerConfig ServerConfig
	GlobalNacosConfig  NacosConfig
)
