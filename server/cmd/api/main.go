// Code generated by hertz generator.

package main

import (
	"NihiStore/server/cmd/api/config"
	"NihiStore/server/cmd/api/initialize"
	"NihiStore/server/cmd/api/initialize/rpc"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
)

func main() {
	initialize.InitLogger()
	r, info := initialize.InitNacos()
	corsCfg := initialize.InitCors()
	tracer, trcCfg := hertztracing.NewServerTracer()
	initialize.InitAliPay()
	rpc.Init()
	h := server.New(
		tracer,
		server.WithALPN(true),
		server.WithHostPorts(fmt.Sprintf(":%d", config.GlobalServerConfig.Port)),
		server.WithRegistry(r, info),
	)
	h.Use(cors.New(corsCfg))
	h.Use(hertztracing.ServerMiddleware(trcCfg))
	register(h)
	h.Spin()
}
