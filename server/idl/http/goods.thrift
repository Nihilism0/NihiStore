namespace go goods
include "../base/common.thrift"
include "../base/goods.thrift"

struct CreateGoodsReq{
    1: i64  id,
    2: goods.Goods goods_information,
}
struct DeleteGoodsReq{
    1: i64 goods_id,
}

struct SearchGoodsInfoReq{
    1: string search_msg
}

struct SearchGoodsReq{
    1: string search_msg
}

service goodsService{
    common.NilResponse CreateGoods(1:CreateGoodsReq req) (api.post ="/goods/create"),
    common.NilResponse DeleteGoods(1:DeleteGoodsReq req) (api.delete ="/goods/delete")
    common.NilResponse SearchGoodsInfo(1:SearchGoodsInfoReq req)(api.get ="/goods/searchgoodsinfo")
    common.NilResponse SearchGoods(1:SearchGoodsReq req)(api.get ="/goods/searchgoods")
}