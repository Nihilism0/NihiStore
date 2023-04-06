namespace go base

struct Goods{
    1: string name,
    2: string describe,
    3: i64 cost,
}

struct GoodsInCart{
    1:i64 amount
    2:i64 goods_id
    3:i64 user_id
}