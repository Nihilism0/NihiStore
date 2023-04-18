package initialize

import (
	"NihiStore/server/cmd/oss/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

func InitMinio() error {
	endpoint := config.GlobalServerConfig.MinioInfo.Endpoint
	accessKeyID := config.GlobalServerConfig.MinioInfo.AccessKeyID
	secretAccessKey := config.GlobalServerConfig.MinioInfo.SecretAccessKey
	useSSL := false
	client, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(
			accessKeyID,
			secretAccessKey,
			""),
		Secure: useSSL,
	})
	if err != nil {
		return err
	}
	if client == nil {
		log.Println("client nil")
	}
	config.MinioClient = client
	return nil
}
