package config

import (
	"NihiStore/server/shared/kitex_gen/goods/goodsservice"
	"NihiStore/server/shared/kitex_gen/user/userservice"
)

var (
	GlobalServerConfig = &ServerConfig{}
	GlobalNacosConfig  = &NacosConfig{}

	GlobalUserClient  userservice.Client
	GlobalGoodsClient goodsservice.Client
)
