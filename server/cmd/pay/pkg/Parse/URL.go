package Parse

import (
	"NihiStore/server/cmd/pay/config"
	"NihiStore/server/shared/kitex_gen/pay"
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"github.com/smartwalle/xid"
	"net/url"
	"strconv"
)

type ParseGenerator struct{}

func (*ParseGenerator) GetTradeNo(userId, GoodsId int64) string {
	var tradeNo = fmt.Sprintf("%d", xid.Next())
	return strconv.FormatInt(userId, 10) + "-" + strconv.FormatInt(GoodsId, 10) + "-" + tradeNo
}

func (*ParseGenerator) GetUrl(in *pay.BuyGoodsRequest, tradeNo, SellerAliId string) *url.URL {
	var p = alipay.TradePagePay{}
	p.NotifyURL = config.GlobalServerConfig.AlipayInfo.KServerDomain + "/alipay/notify"
	p.ReturnURL = config.GlobalServerConfig.AlipayInfo.KServerDomain + "/alipay/callback"
	p.Subject = in.Subject + tradeNo
	p.OutTradeNo = tradeNo
	p.TotalAmount = in.TotalAmount
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	if SellerAliId != "" {
		p.SellerId = SellerAliId
	}
	URL, _ := config.AliClient.TradePagePay(p)
	return URL
}
