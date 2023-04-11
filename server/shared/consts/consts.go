package consts

const (
	JWTIssuer  = "NihiStore"
	ThirtyDays = 60 * 60 * 24 * 30

	IPFlagName  = "ip"
	IPFlagValue = "0.0.0.0"
	IPFlagUsage = "address"

	PortFlagName  = "port"
	PortFlagUsage = "port"

	TCP             = "tcp"
	FreePortAddress = "localhost:0"

	HlogFilePath = "./tmp/hlog/logs/"
	KlogFilePath = "./tmp/klog/logs/"

	ApiConfigPath   = "config.yaml"
	UserConfigPath  = "config.yaml"
	GoodsConfigPath = "config.yaml"
	PayConfigPath   = "config.yaml"

	NacosLogDir   = "tmp/nacos/log"
	NacosCacheDir = "tmp/nacos/cache"
	NacosLogLevel = "debug"

	ApiGroup   = "API_GROUP"
	UserGroup  = "USER_GROUP"
	GoodsGroup = "GOODS_GROUP"

	PayGroup = "PAY_GROUP"
	MySqlDSN = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	MaxGoodsAmount = 99
)
