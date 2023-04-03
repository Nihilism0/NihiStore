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
	UserSrvInfo  RPCSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	GoodsSrvInfo RPCSrvConfig `mapstructure:"goods_srv" json:"goods_srv"`
}

type RPCSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
