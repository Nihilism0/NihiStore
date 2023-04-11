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

//更新商品
struct UpdateGoodsRequest{
    1: i64  id,
    2: goods.Goods goods_information,
    3: i64 userId
}

struct UpdateGoodsResponse{
    1: common.BaseResponse base_resp
}

//搜索局部信息
struct SearchGoodsInfoRequest{
        1: string searchMsg
        2: i64 page
        3: i64 pageAmount
}

struct SearchGoodsInfoResponse{
    1: common.BaseResponse base_resp
    2: list<goods.Name> names
}

//得到完全信息
struct SearchGoodsRequest{
    1: i64 id
}

struct SearchGoodsResponse{
    1: goods.GoodsFullInfo goodsFI
    2: common.BaseResponse base_resp
}

service GoodsService {
    CreateGoodsResponse CreateGoods(1:CreateGoodsRequest req),
    DeleteGoodsResponse DeleteGoods(1:DeleteGoodsRequest req)
    UpdateGoodsResponse UpdateGoods(1:UpdateGoodsRequest req)
    SearchGoodsInfoResponse SearchGoodsInfo(1:SearchGoodsInfoRequest req),
    SearchGoodsResponse SearchGoods(1:SearchGoodsRequest req),
}