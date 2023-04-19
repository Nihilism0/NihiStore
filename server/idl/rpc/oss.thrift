namespace go oss

include "../base/common.thrift"

struct CreateGoodsOSSRequest{
    1: string path,
    2: i32 timeout_sec,
    3: i64 goodsId
}

struct CreateGoodsOSSResponse{
    1: common.BaseResponse base_resp,
    2: string uploadUrl,
    3: i64 id,
}

struct GetGoodsOSSRequest{
    1: i64 id,
    2: i32 timeout_sec
}

struct GetGoodsOSSResponse{
    1: common.BaseResponse base_resp,
    2: string url
}

struct CreateHeadOSSRequest{
    1: string path,
    2: i32 timeout_sec,
    3: i64 userId
}

struct CreateHeadOSSResponse{
    1: common.BaseResponse base_resp,
    2: string uploadUrl,
    3: i64 id,
}

struct GetHeadOSSRequest{
    1: i64 id,
    2: i32 timeout_sec
}

struct GetHeadOSSResponse{
    1: common.BaseResponse base_resp,
    2: string url
}


service OSSService{
    CreateGoodsOSSResponse CreateGoodsOSS(1: CreateGoodsOSSRequest req),
    GetGoodsOSSResponse GetGoodsOSS(1: GetGoodsOSSRequest req),
    CreateHeadOSSResponse CreateHeadOSS(1: CreateHeadOSSRequest req),
    GetHeadOSSResponse GetHeadOSS(1: GetHeadOSSRequest req),
}