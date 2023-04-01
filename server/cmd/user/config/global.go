package config

import "gorm.io/gorm"

var (
	GlobalServerConfig ServerConfig
	GlobalConsulConfig ConsulConfig
	DB                 *gorm.DB
)
