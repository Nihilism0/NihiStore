namespace go base

struct Goods{
    1: string name,
    2: string describe,
    3: i64 cost,
}

struct GoodsFullInfo{
    1: i64 id
    2: string name
    3: i64 user_id
    4: string describe
    5: i64 cost
    6: i64 sales_volume
}

struct Name{
    1: string name,
    2: i64 id,
}

struct GoodsInCart{
    1: i64 amount
    2: i64 goods_id
    3: i64 user_id
}