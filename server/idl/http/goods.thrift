namespace go goods
include "../base/common.thrift"
include "../base/goods.thrift"

struct CreateGoodsRequest{
    1: i64  id,
    2: goods.Goods goods_information,
}

struct CreateGoodsResponse{
    1:common.BaseResponse base_resp,
}

service goodsService{
    common.NilResponse CreateGoods(1:CreateGoodsRequest req) (api.post ="/goods/create")
}