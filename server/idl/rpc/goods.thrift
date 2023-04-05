include "../base/common.thrift"
include "../base/goods.thrift"

//创建商品
struct CreateGoodsRequest{
    1: i64  id,
    2: goods.Goods goods_information,
}

struct CreateGoodsResponse{
    1:common.BaseResponse base_resp,
}

//删除商品
struct DeleteGoodsRequest{
    1: i64 goods_id
    2: i64 seller_id
}

struct DeleteGoodsResponse{
    1:common.BaseResponse base_resp
}

//搜索局部信息
struct SearchGoodsInfoRequest{
    1: string search_msg
}

struct SearchGoodsInfoResponse{
    1: string name
}

//搜索完全信息
struct SearchGoodsRequest{
    1: string search_msg
}

struct SearchGoodsResponse{
    1: goods.Goods goods_information
    2: i64 sales_volume
}

service GoodsService {
    CreateGoodsResponse CreateGoods(1:CreateGoodsRequest req),
    DeleteGoodsResponse DeleteGoods(1:DeleteGoodsRequest req)
    SearchGoodsInfoResponse SearchGoodsInfo(1:SearchGoodsInfoRequest req),
    SearchGoodsResponse SearchGoods(1:SearchGoodsRequest req),
}