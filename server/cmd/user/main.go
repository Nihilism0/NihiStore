package main

import (
	"NihiStore/server/cmd/user/config"
	"NihiStore/server/cmd/user/initialize"
	"NihiStore/server/cmd/user/pkg/jwt"
	"NihiStore/server/cmd/user/pkg/mysql"
	"NihiStore/server/shared/consts"
	user "NihiStore/server/shared/kitex_gen/user/userservice"
	"NihiStore/server/shared/middleware"
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
	// Create new server.
	srv := user.NewServer(
		&UserServiceImpl{
			TokenGenerator:     &jwt.TokenGenerator{},
			MysqlUserGenerator: &mysql.MysqlUserGenerator{},
			MysqlFavoGenerator: &mysql.MysqlFavoGenerator{},
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
