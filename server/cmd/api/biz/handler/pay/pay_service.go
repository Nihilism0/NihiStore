// Code generated by hertz generator.

package pay

import (
	base "NihiStore/server/cmd/api/biz/model/base"
	hpay "NihiStore/server/cmd/api/biz/model/pay"
	"NihiStore/server/cmd/api/config"
	kpay "NihiStore/server/shared/kitex_gen/pay"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/smartwalle/alipay/v3"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// BuyGoods .
// @router /alipay/pay [GET]
func BuyGoods(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hpay.BuyGoodsReq
	resp := new(kpay.BuyGoodsResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	//ID, _ := c.Get("ID")
	resp, err = config.GlobalPayClient.BuyGoods(ctx, &kpay.BuyGoodsRequest{
		Subject:     req.Subject,
		TotalAmount: req.TotalAmount,
		GoodsId:     req.GoodsId,
		UserId:      66,
	})
	if err != nil {
		hlog.Error("rpc user service err!", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	fmt.Println(resp.URL)
	c.JSON(http.StatusOK, resp.BaseResp)
	c.Redirect(http.StatusTemporaryRedirect, []byte(resp.URL))
}

// CallBack .
// @router /alipay/callback [GET]
func CallBack(ctx context.Context, c *app.RequestContext) {
	var err error
	k, _ := c.MultipartForm()
	ok, err := config.AliClient.VerifySign(k.Value)
	if err != nil {
		hlog.Error("回调验证签名发生错误", err.Error())
		c.String(http.StatusBadRequest, "回调验证签名发生错误")
		return
	}

	if ok == false {
		hlog.Error("回调验证签名未通过")
		c.String(http.StatusBadRequest, "回调验证签名未通过")
		return
	}
	hlog.Info("回调验证签名通过")

	var outTradeNo, _ = c.GetPostForm("out_trade_no")
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo
	rsp, err := config.AliClient.TradeQuery(p)
	if err != nil {
		c.String(http.StatusBadRequest, "验证订单 %s 信息发生错误: %s", outTradeNo, err.Error())
		return
	}
	if rsp.IsSuccess() == false {
		c.String(http.StatusBadRequest, "验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Content.Msg, rsp.Content.SubMsg)
		return
	}
	c.String(http.StatusOK, "订单 %s 支付成功", outTradeNo)
}

// Notify .
// @router /alipay/notify [POST]
func Notify(ctx context.Context, c *app.RequestContext) {

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}
