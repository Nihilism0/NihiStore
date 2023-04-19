namespace go user
include "../base/common.thrift"

struct LoginReq {
    1: string Username (api.form="username" api.vd = "len($) > 0 && len($) < 20")//名字长度大于0小于20
    2: string Password (api.form="password" api.vd = "len($) > 0 && len($) < 20")//密码长度大于0小于20
}

struct RegisterReq {
    1: string Username (api.form="username" api.vd = "len($) > 0 && len($) < 20")//名字长度大于0小于20
    2: string Password (api.form="password" api.vd = "len($) > 0 && len($) < 20")//密码长度大于0小于20
}

struct CreateFavoritesReq {
    1: string favoritesName (api.form="favoritesname" api.vd = "len($) > 0 && len($) < 20")
    2: string description (api.form="description" api.vd = "len($) > 0 && len($) < 20")
}

struct DeleteFavoritesReq{
    1: i64 favoritesId (api.form="favoritesId" )
}

struct CollectGoodsReq{
    1: i64 goodsId
    2: i64 favoritesId
}

struct RmCollectGoodsReq{
    1: i64 goodsId
    2: i64 favoritesId
    3: i64 userId
}

struct WatchFavoritesReq{}

struct WatchGoodsInFavoritesReq{
    1: i64 favoritesId
}

struct AddToCartReq{
    1: i64 goodsId (api.form="goodsId")
    2: i64 amount (api.form="amount" api.vd = "$ > 0 && $ < 100")
}

struct AddAmountCartReq{
    1:i64 goodsId
    2:i64 amount
}

struct DeleteAmountCartReq{
    1:i64 goodsId
    2:i64 amount
}

struct RemoveOutCartReq{
    1: i64 goodsId
}

struct WatchCartReq{}

struct CleanCartReq{}

struct BeSellerReq{
    1: string sellerAliId
}

struct GetSellerByGoodsReq{
    1: string goodsId
}

struct UploadHeadReq{}

struct GetHeadReq{}

service userService {
    common.NilResponse Login(1: LoginReq request) (api.get="/user/login")
    common.NilResponse Register(1: RegisterReq request) (api.post="/user/register")
    common.NilResponse CreateFavorites(1: CreateFavoritesReq request) (api.post="/user/favorites/createfavorites")
    common.NilResponse WatchFavorites(1: WatchFavoritesReq request) (api.get="/user/favorites/watchfavorites")
    common.NilResponse WatchGoodsInFavorites(1: WatchGoodsInFavoritesReq request) (api.get="/user/favorites/watchgoodsinfavorites")
    common.NilResponse DeleteFavorites(1: DeleteFavoritesReq request) (api.delete="/user/favorites/deletefavorites")
    common.NilResponse CollectGoods(1: CollectGoodsReq request) (api.post="/user/favorites/collectgoods")
    common.NilResponse AddToCart(1: AddToCartReq request) (api.post="/user/cart/addtocart")
    common.NilResponse AddAmounrCart(1: AddAmountCartReq request) (api.put="/user/cart/addamountcart")
    common.NilResponse DeleteAmountCart(1: DeleteAmountCartReq request) (api.put="/user/cart/deleteamountcart")
    common.NilResponse RemoveOutCart(1: RemoveOutCartReq request) (api.delete="/user/cart/removeoutcart")
    common.NilResponse WatchCart(1: WatchCartReq request) (api.get="/user/cart/watchcart")
    common.NilResponse CleanCart(1: CleanCartReq request) (api.delete="/user/cart/cleancart")
    common.NilResponse BeSeller(1:BeSellerReq request) (api.put="/user/beseller")
    common.NilResponse GetSellerByGoods(1:GetSellerByGoodsReq request) (api.get="/user/getsellerbygoods")
    common.NilResponse UploadHead(1:UploadHeadReq request) (api.post="/user/uploadhead")
    common.NilResponse GetHead(1:GetHeadReq request) (api.get="/user/gethead")
}