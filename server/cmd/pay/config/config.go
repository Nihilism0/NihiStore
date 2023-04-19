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

type ServerConfig struct {
	Name       string       `mapstructure:"name" json:"name"`
	Host       string       `mapstructure:"host" json:"host"`
	OtelInfo   OtelConfig   `mapstructure:"otel" json:"otel"`
	MysqlInfo  MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	JWTInfo    JWTConfig    `mapstructure:"jwt" json:"jwt"`
	MqInfo     MqConfig     `mapstructure:"mq_info" json:"mq_info"`
	AlipayInfo AlipayConfig `mapstructure:"alipay" json:"alipay"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
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

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}
