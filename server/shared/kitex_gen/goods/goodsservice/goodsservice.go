// Code generated by Kitex v0.4.4. DO NOT EDIT.

package goodsservice

import (
	goods "NihiStore/server/shared/kitex_gen/goods"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return goodsServiceServiceInfo
}

var goodsServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "GoodsService"
	handlerType := (*goods.GoodsService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateGoods": kitex.NewMethodInfo(createGoodsHandler, newGoodsServiceCreateGoodsArgs, newGoodsServiceCreateGoodsResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "goods",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func createGoodsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*goods.GoodsServiceCreateGoodsArgs)
	realResult := result.(*goods.GoodsServiceCreateGoodsResult)
	success, err := handler.(goods.GoodsService).CreateGoods(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newGoodsServiceCreateGoodsArgs() interface{} {
	return goods.NewGoodsServiceCreateGoodsArgs()
}

func newGoodsServiceCreateGoodsResult() interface{} {
	return goods.NewGoodsServiceCreateGoodsResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateGoods(ctx context.Context, req *goods.CreateGoodsRequest) (r *goods.CreateGoodsResponse, err error) {
	var _args goods.GoodsServiceCreateGoodsArgs
	_args.Req = req
	var _result goods.GoodsServiceCreateGoodsResult
	if err = p.c.Call(ctx, "CreateGoods", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
