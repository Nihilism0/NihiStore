package config

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type ServerConfig struct {
	Name        string       `mapstructure:"name" json:"name"`
	Host        string       `mapstructure:"host" json:"host"`
	Port        int          `mapstructure:"port" json:"port"`
	JWTInfo     JWTConfig    `mapstructure:"jwt" json:"jwt"`
	UserSrvInfo RPCSrvConfig `mapstructure:"user_srv" json:"user_srv"`
}

type RPCSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
