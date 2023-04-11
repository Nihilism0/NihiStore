package initialize

import (
	"NihiStore/server/cmd/api/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/smartwalle/alipay/v3"
	"os"
	"path/filepath"
)

func InitAliPay() {
	var err error

	if config.AliClient, err = alipay.New(config.GlobalServerConfig.AlipayInfo.KAppId, config.GlobalServerConfig.AlipayInfo.KPrivateKey, false); err != nil {
		hlog.Fatalf("初始化支付宝失败", err.Error())
		return
	}
	// 加载证书
	dir, _ := os.Getwd()
	certPath := filepath.Join(dir, "cert/appPublicCert.crt")
	if err = config.AliClient.LoadAppPublicCertFromFile(certPath); err != nil {
		hlog.Fatalf("加载证书发生错误", err.Error())
		return
	}
	certPath = filepath.Join(dir, "cert/alipayRootCert.crt")
	if err = config.AliClient.LoadAliPayRootCertFromFile(certPath); err != nil {
		hlog.Fatalf("加载证书发生错误", err.Error())
		return
	}
	certPath = filepath.Join(dir, "cert/alipayPublicCert.crt")
	if err = config.AliClient.LoadAliPayPublicCertFromFile(certPath); err != nil {
		hlog.Fatalf("加载证书发生错误", err.Error())
		return
	}
	if err = config.AliClient.SetEncryptKey(config.GlobalServerConfig.AlipayInfo.EncryptKey); err != nil {
		hlog.Fatalf("加载内容加密密钥发生错误", err.Error())
		return
	}
}
