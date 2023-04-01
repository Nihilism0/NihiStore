namespace go base

struct BaseResponse {
    1: i64 status_code,
    2: string status_msg,
}

struct NilResponse {}