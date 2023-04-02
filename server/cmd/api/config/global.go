package config

import (
	"NihiStore/server/shared/kitex_gen/user/userservice"
)

var (
	GlobalServerConfig = &ServerConfig{}
	GlobalNacosConfig  = &NacosConfig{}

	GlobalUserClient userservice.Client
)
