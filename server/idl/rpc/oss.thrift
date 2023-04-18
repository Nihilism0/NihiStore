namespace go oss

include "../base/common.thrift"

struct CreateOSSRequest{
    1: string path,
    2: i32 timeout_sec,
}

struct CreateOSSResponse{
    1: common.BaseResponse base_resp,
    2: string uploadUrl,
    3: string id,
}

struct GetOSSRequest{
    1: string path,
    2: i32 timeout_sec
}

struct GetOSSResponse{
    1: common.BaseResponse base_resp,
    2: string url
}

service OSSService{
    CreateOSSResponse CreateOSS(1: CreateOSSRequest req),
    GetOSSResponse GetOSS(1: GetOSSRequest req),
}