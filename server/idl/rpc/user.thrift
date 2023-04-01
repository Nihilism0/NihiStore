include "../base/common.thrift"
struct LoginRequest {
    1: string username,
}

struct LoginResponse {
    1: i64 account_id,
}

struct RegisterRequest{
    1: string username,
    2: string password,
}

struct RegisterResponse {
    1: common.BaseResponse base_resp,
}

service UserService {
    LoginResponse Login(1:LoginRequest req),
    RegisterResponse Register(1:RegisterRequest req),
}
