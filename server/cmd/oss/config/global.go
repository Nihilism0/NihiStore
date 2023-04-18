package config

import (
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

var (
	DB                 *gorm.DB
	GlobalServerConfig ServerConfig
	GlobalNacosConfig  NacosConfig
	MinioClient        *minio.Client
)
