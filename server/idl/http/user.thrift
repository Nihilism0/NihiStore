namespace go user
include "../base/common.thrift"

struct LoginReq {
    1: string Name (api.form="name");
    2: string Password (api.form="password");
}

struct LoginResp {
    1: common.BaseResponse base_resp
    2: string token;
}

struct RegisterReq {
    1: string Name (api.form="name");
    2: string Pssword (api.form="password");
}

struct RegisterResp {
    1: common.BaseResponse base_resp
    2: string Msg;
}

service userService {
    LoginResp Login(1: LoginReq request) (api.get="/user/login")
    RegisterResp Register(1: RegisterReq request) (api.get="/user/request")
}