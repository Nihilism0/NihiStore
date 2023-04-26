package config

import (
	"github.com/casbin/casbin/v2"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

var (
	DB                 *gorm.DB
	MinioClient        *minio.Client
	GlobalServerConfig ServerConfig
	GlobalNacosConfig  NacosConfig
	Enforcer           *casbin.Enforcer
)
