namespace go pay
include "../base/common.thrift"
//购买商品
struct BuyGoodsReq{
    1: string Subject
    2: i64 goodsId
    3: string TotalAmount
}

struct CallBackReq{}
struct NotifyReq{}

service payService {
    common.NilResponse BuyGoods(1: BuyGoodsReq request) (api.get="/alipay/pay")
    common.NilResponse CallBack(1: CallBackReq request) (api.get="/alipay/callback")
    common.NilResponse Notify(1: NotifyReq request) (api.post="/alipay/notify")
}