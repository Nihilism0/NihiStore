namespace go goods
include "../base/common.thrift"
include "../base/goods.thrift"

struct CreateGoodsReq{
    1: i64  id,
    2: goods.Goods goods_information,
}
struct DeleteGoodsReq{
    1: i64 goodsId,
}

struct UpdateGoodsReq{
    1: i64  id,
    2: goods.Goods goods_information,
}

struct SearchGoodsInfoReq{
    1: string searchMsg
    2: i64 page
    3: i64 pageAmount
}

struct SearchGoodsReq{
    1: i64 goodsId
}

service goodsService{
    common.NilResponse CreateGoods(1:CreateGoodsReq req) (api.post ="/goods/create"),
    common.NilResponse DeleteGoods(1:DeleteGoodsReq req) (api.delete ="/goods/delete")
    common.NilResponse UpdateGoods(1:UpdateGoodsReq req) (api.put="/goods/update")
    common.NilResponse SearchGoodsInfo(1:SearchGoodsInfoReq req)(api.get ="/goods/searchgoodsinfo")
    common.NilResponse SearchGoods(1:SearchGoodsReq req)(api.get ="/goods/searchgoods")
}