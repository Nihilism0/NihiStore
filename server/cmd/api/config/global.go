package config

import (
	"NihiStore/server/shared/kitex_gen/goods/goodsservice"
	pay "NihiStore/server/shared/kitex_gen/pay/payservice"
	"NihiStore/server/shared/kitex_gen/user/userservice"
	"github.com/smartwalle/alipay/v3"
)

var (
	GlobalServerConfig = &ServerConfig{}
	GlobalNacosConfig  = &NacosConfig{}
	AliClient          *alipay.Client
	GlobalUserClient   userservice.Client
	GlobalGoodsClient  goodsservice.Client
	GlobalPayClient    pay.Client
)
