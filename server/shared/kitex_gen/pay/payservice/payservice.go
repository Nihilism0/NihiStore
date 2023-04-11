// Code generated by Kitex v0.5.1. DO NOT EDIT.

package payservice

import (
	pay "NihiStore/server/shared/kitex_gen/pay"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return payServiceServiceInfo
}

var payServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "PayService"
	handlerType := (*pay.PayService)(nil)
	methods := map[string]kitex.MethodInfo{
		"BuyGoods": kitex.NewMethodInfo(buyGoodsHandler, newPayServiceBuyGoodsArgs, newPayServiceBuyGoodsResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "pay",
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

func buyGoodsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pay.PayServiceBuyGoodsArgs)
	realResult := result.(*pay.PayServiceBuyGoodsResult)
	success, err := handler.(pay.PayService).BuyGoods(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPayServiceBuyGoodsArgs() interface{} {
	return pay.NewPayServiceBuyGoodsArgs()
}

func newPayServiceBuyGoodsResult() interface{} {
	return pay.NewPayServiceBuyGoodsResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) BuyGoods(ctx context.Context, req *pay.BuyGoodsRequest) (r *pay.BuyGoodsResponse, err error) {
	var _args pay.PayServiceBuyGoodsArgs
	_args.Req = req
	var _result pay.PayServiceBuyGoodsResult
	if err = p.c.Call(ctx, "BuyGoods", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
