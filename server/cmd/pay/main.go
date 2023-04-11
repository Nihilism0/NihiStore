package main

import (
	"NihiStore/server/cmd/pay/config"
	"NihiStore/server/cmd/pay/initialize"
	"NihiStore/server/cmd/pay/pkg/Parse"
	"NihiStore/server/shared/consts"
	pay "NihiStore/server/shared/kitex_gen/pay/payservice"
	"NihiStore/server/shared/middleware"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"net"
	"strconv"
)

func main() {
	initialize.InitLogger()
	IP, Port := initialize.InitFlag()
	r, info := initialize.InitNacos(Port)
	initialize.InitDB()
	initialize.InitAliPay()
	// Create new server.
	srv := pay.NewServer(
		&PayServiceImpl{
			ParseGenerator: &Parse.ParseGenerator{},
		},
		server.WithServiceAddr(utils.NewNetAddr(consts.TCP, net.JoinHostPort(IP, strconv.Itoa(Port)))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}),
	)
	fmt.Println(config.GlobalServerConfig.AlipayInfo.KAppId)
	err := srv.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
