// Code generated by Kitex v0.5.1. DO NOT EDIT.

package ossservice

import (
	oss "NihiStore/server/shared/kitex_gen/oss"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return oSSServiceServiceInfo
}

var oSSServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "OSSService"
	handlerType := (*oss.OSSService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateGoodsOSS": kitex.NewMethodInfo(createGoodsOSSHandler, newOSSServiceCreateGoodsOSSArgs, newOSSServiceCreateGoodsOSSResult, false),
		"GetGoodsOSS":    kitex.NewMethodInfo(getGoodsOSSHandler, newOSSServiceGetGoodsOSSArgs, newOSSServiceGetGoodsOSSResult, false),
		"CreateHeadOSS":  kitex.NewMethodInfo(createHeadOSSHandler, newOSSServiceCreateHeadOSSArgs, newOSSServiceCreateHeadOSSResult, false),
		"GetHeadOSS":     kitex.NewMethodInfo(getHeadOSSHandler, newOSSServiceGetHeadOSSArgs, newOSSServiceGetHeadOSSResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "oss",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.1",
		Extra:           extra,
	}
	return svcInfo
}

func createGoodsOSSHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*oss.OSSServiceCreateGoodsOSSArgs)
	realResult := result.(*oss.OSSServiceCreateGoodsOSSResult)
	success, err := handler.(oss.OSSService).CreateGoodsOSS(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOSSServiceCreateGoodsOSSArgs() interface{} {
	return oss.NewOSSServiceCreateGoodsOSSArgs()
}

func newOSSServiceCreateGoodsOSSResult() interface{} {
	return oss.NewOSSServiceCreateGoodsOSSResult()
}

func getGoodsOSSHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*oss.OSSServiceGetGoodsOSSArgs)
	realResult := result.(*oss.OSSServiceGetGoodsOSSResult)
	success, err := handler.(oss.OSSService).GetGoodsOSS(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOSSServiceGetGoodsOSSArgs() interface{} {
	return oss.NewOSSServiceGetGoodsOSSArgs()
}

func newOSSServiceGetGoodsOSSResult() interface{} {
	return oss.NewOSSServiceGetGoodsOSSResult()
}

func createHeadOSSHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*oss.OSSServiceCreateHeadOSSArgs)
	realResult := result.(*oss.OSSServiceCreateHeadOSSResult)
	success, err := handler.(oss.OSSService).CreateHeadOSS(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOSSServiceCreateHeadOSSArgs() interface{} {
	return oss.NewOSSServiceCreateHeadOSSArgs()
}

func newOSSServiceCreateHeadOSSResult() interface{} {
	return oss.NewOSSServiceCreateHeadOSSResult()
}

func getHeadOSSHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*oss.OSSServiceGetHeadOSSArgs)
	realResult := result.(*oss.OSSServiceGetHeadOSSResult)
	success, err := handler.(oss.OSSService).GetHeadOSS(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOSSServiceGetHeadOSSArgs() interface{} {
	return oss.NewOSSServiceGetHeadOSSArgs()
}

func newOSSServiceGetHeadOSSResult() interface{} {
	return oss.NewOSSServiceGetHeadOSSResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateGoodsOSS(ctx context.Context, req *oss.CreateGoodsOSSRequest) (r *oss.CreateGoodsOSSResponse, err error) {
	var _args oss.OSSServiceCreateGoodsOSSArgs
	_args.Req = req
	var _result oss.OSSServiceCreateGoodsOSSResult
	if err = p.c.Call(ctx, "CreateGoodsOSS", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetGoodsOSS(ctx context.Context, req *oss.GetGoodsOSSRequest) (r *oss.GetGoodsOSSResponse, err error) {
	var _args oss.OSSServiceGetGoodsOSSArgs
	_args.Req = req
	var _result oss.OSSServiceGetGoodsOSSResult
	if err = p.c.Call(ctx, "GetGoodsOSS", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateHeadOSS(ctx context.Context, req *oss.CreateHeadOSSRequest) (r *oss.CreateHeadOSSResponse, err error) {
	var _args oss.OSSServiceCreateHeadOSSArgs
	_args.Req = req
	var _result oss.OSSServiceCreateHeadOSSResult
	if err = p.c.Call(ctx, "CreateHeadOSS", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetHeadOSS(ctx context.Context, req *oss.GetHeadOSSRequest) (r *oss.GetHeadOSSResponse, err error) {
	var _args oss.OSSServiceGetHeadOSSArgs
	_args.Req = req
	var _result oss.OSSServiceGetHeadOSSResult
	if err = p.c.Call(ctx, "GetHeadOSS", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
