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

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Salt     string `mapstructure:"salt" json:"salt"`
}

type ServerConfig struct {
	Name       string       `mapstructure:"name" json:"name"`
	Host       string       `mapstructure:"host" json:"host"`
	OtelInfo   OtelConfig   `mapstructure:"otel" json:"otel"`
	MysqlInfo  MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	JWTInfo    JWTConfig    `mapstructure:"jwt" json:"jwt"`
	OSSSrvInfo OSSSrvConfig `mapstructure:"oss_srv" json:"oss_srv"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type OSSSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}
