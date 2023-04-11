include "../base/common.thrift"

struct BuyGoodsRequest{
     1: string Subject
     2: i64 userId
     3: string TotalAmount
     4: i64 goodsId
}

struct BuyGoodsResponse {
    1: common.BaseResponse base_resp
    2: string URL
}

service PayService{
    BuyGoodsResponse BuyGoods(1: BuyGoodsRequest req)
}