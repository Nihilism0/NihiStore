package main

import (
	"NihiStore/server/cmd/goods/config"
	"NihiStore/server/cmd/goods/initialize"
	"NihiStore/server/cmd/goods/pkg/convert"
	"NihiStore/server/cmd/goods/pkg/mysql"
	"NihiStore/server/shared/consts"
	goods "NihiStore/server/shared/kitex_gen/goods/goodsservice"
	"NihiStore/server/shared/middleware"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"gorm.io/plugin/opentelemetry/provider"
	"net"
	"strconv"
)

func main() {
	initialize.InitLogger()
	IP, Port := initialize.InitFlag()
	r, info := initialize.InitNacos(Port)
	initialize.InitDB()
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
	)
	defer p.Shutdown(context.Background())
	ossClient := initialize.InitOSS()
	// Create new server.
	srv := goods.NewServer(&GoodsServiceImpl{
		ConvertGenerator: &convert.ConvertGenerator{},
		MysqlGenerator:   &mysql.MysqlGenerator{},
		OSSManager:       ossClient,
	},
		server.WithServiceAddr(utils.NewNetAddr(consts.TCP, net.JoinHostPort(IP, strconv.Itoa(Port)))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}),
	)

	err := srv.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
