include "../base/common.thrift"
struct LoginRequest {
    1: string username,
    2: string password,
}

struct LoginResponse {
    1: common.BaseResponse base_resp
    2: string token,
}

struct RegisterRequest{
    1: string username,
    2: string password,
}

struct RegisterResponse {
    1: common.BaseResponse base_resp,
    2: bool OK
}

service UserService {
    LoginResponse Login(1:LoginRequest req),
    RegisterResponse Register(1:RegisterRequest req),
}
