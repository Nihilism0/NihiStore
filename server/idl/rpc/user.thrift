include "../base/common.thrift"
include "../base/goods.thrift"
include "../base/favorites.thrift"

//登陆
struct LoginRequest {
    1: string username
    2: string password
}

struct LoginResponse {
    1: common.BaseResponse base_resp
    2: string token
}

//注册
struct RegisterRequest{
    1: string username
    2: string password
}

struct RegisterResponse {
    1: common.BaseResponse base_resp,
}

//创建收藏夹
struct CreateFavoritesRequest {
   1: string favoritesName
   2: string description
   3: i64 userId
}

struct CreateFavoritesResponse{
    1:common.BaseResponse base_resp
}

//看收藏夹
struct WatchFavoritesRequset{
    1: i64 userId
}

struct WatchFavoritesResponse{
    1: common.BaseResponse base_resp
    2: list<favorites.Favorites> favoriteses
}

//看收藏夹里的商品
struct WatchGoodsInFavoritesRequest{
    1: i64 favoritesId
    2: i64 userId
}

struct WatchGoodsInFavoritesResponse{
    1: common.BaseResponse base_resp
    2: list <favorites.GoodsInFavorites> goodsIds
}

//删除收藏夹
struct DeleteFavoritesRequest{
    1:i64 favoritesId
    2:i64 userId
}

struct DeleteFavoritesResponse{
    1:common.BaseResponse base_resp
}

//收藏商品
struct CollectGoodsRequest{
    1: i64 goodsId
    2: i64 favoritesId
    3: i64 userId
}

struct CollectGoodsResponse{
    1: common.BaseResponse base_resp
}

//取消收藏商品
struct RmCollectGoodsRequest{
    1: i64 goodsId
    2: i64 favoritesId
    3: i64 userId
}

struct RmCollectGoodsResponse{
    1: common.BaseResponse base_resp
}

//添加到购物车
struct AddToCartRequest{
    1: i64 goodsId
    2: i64 amount
    3: i64 userId
}

struct AddToCartResponse{
    1:common.BaseResponse base_resp
}

//增加购物车商品数量
struct AddAmountCartRequest{
    1:i64 goodsId
    2:i64 amount
    3:i64 userId
}

struct AddAmountCartResponse{
    1:common.BaseResponse base_resp
}
//减少购物车商品数量
struct DeleteAmountCartRequest{
    1:i64 goodsId
    2:i64 amount
    3:i64 userId
}

struct DeleteAmountCartResponse{
    1:common.BaseResponse base_resp
}

//删除购物车中的商品
struct RemoveOutCartRequest{
    1: i64 goodsId
    2: i64 userId
}

struct RemoveOutCartResponse{
    1:common.BaseResponse base_resp
}

//查看购物车
struct WatchCartRequest{
    1: i64 userId
}

struct WatchCartResponse{
    1: common.BaseResponse base_resp
    2: list<goods.GoodsInCart> goods
}

//清空购物车
struct CleanCartRequest {
    1: i64 userId
}

struct CleanCartResponse {
    1: common.BaseResponse base_resp
}



service UserService {
    LoginResponse Login(1:LoginRequest req)
    RegisterResponse Register(1:RegisterRequest req)
    CreateFavoritesResponse CreateFavorites(1:CreateFavoritesRequest req)
    WatchFavoritesResponse WatchFavorites(1:WatchFavoritesRequset req)
    WatchGoodsInFavoritesResponse WatchGoodsInFavorites(1:WatchGoodsInFavoritesRequest req)
    DeleteFavoritesResponse DeleteFavorites(1:DeleteFavoritesRequest req)
    CollectGoodsResponse CollectGoods(1:CollectGoodsRequest req)
    AddToCartResponse AddToCart(1:AddToCartRequest req)
    AddAmountCartResponse AddAmountCart(1:AddAmountCartRequest req)
    DeleteAmountCartResponse DeleteAmountCart(1:DeleteAmountCartRequest req)
    RemoveOutCartResponse RemoveOutCart(1:RemoveOutCartRequest req)
    WatchCartResponse WatchCart(1:WatchCartRequest req)
    CleanCartResponse CleanCart(1:CleanCartRequest req)
}
