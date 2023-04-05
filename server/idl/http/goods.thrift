namespace go goods
include "../base/common.thrift"
include "../base/goods.thrift"

struct CreateGoodsRequest{
    1: i64  id,
    2: goods.Goods goods_information,
}
struct DeleteGoodsRequest{
    1: i64 goods_id,
}

struct SearchGoodsInfoRequest{
    1: string search_msg
}

struct SearchGoodsRequest{
    1: string search_msg
}

service goodsService{
    common.NilResponse CreateGoods(1:CreateGoodsRequest req) (api.post ="/goods/create"),
    common.NilResponse DeleteGoods(1:DeleteGoodsRequest req) (api.delete ="/goods/delete")
    common.NilResponse SearchGoodsInfo(1:SearchGoodsInfoRequest req)(api.get ="/goods/searchgoodsinfo")
    common.NilResponse SearchGoods(1:SearchGoodsRequest req)(api.get ="/goods/searchgoods")
}