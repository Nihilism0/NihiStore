package config

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type ServerConfig struct {
	Name         string       `mapstructure:"name" json:"name"`
	Host         string       `mapstructure:"host" json:"host"`
	Port         int          `mapstructure:"port" json:"port"`
	JWTInfo      JWTConfig    `mapstructure:"jwt" json:"jwt"`
	MqInfo       MqConfig     `mapstructure:"mq_info" json:"mq_info"`
	OtelInfo     OtelConfig   `mapstructure:"otel" json:"otel"`
	AlipayInfo   AlipayConfig `mapstructure:"alipay" json:"alipay"`
	UserSrvInfo  RPCSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	GoodsSrvInfo RPCSrvConfig `mapstructure:"goods_srv" json:"goods_srv"`
	PaySrvInfo   RPCSrvConfig `mapstructure:"pay_srv" json:"pay_srv"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type RPCSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type AlipayConfig struct {
	KAppId        string `mapstructure:"kAppId" json:"kAppId"`
	KPrivateKey   string `mapstructure:"kPrivateKey" json:"kPrivateKey"`
	KServerPort   string `mapstructure:"kServerPort" json:"kServerPort"`
	KServerDomain string `mapstructure:"kServerDomain" json:"kServerDomain"`
	EncryptKey    string `mapstructure:"EncryptKey" json:"EncryptKey"`
}

type MqConfig struct {
	Address string `mapstructure:"address" json:"address"`
}
