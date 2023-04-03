namespace go user
include "../base/common.thrift"

struct LoginReq {
    1: string Username (api.form="username" api.vd = "len($) > 0 && len($) < 20");//名字长度大于0小于20
    2: string Password (api.form="password" api.vd = "len($) > 0 && len($) < 20");//密码长度大于0小于20
}

struct RegisterReq {
    1: string Username (api.form="username" api.vd = "len($) > 0 && len($) < 20");//名字长度大于0小于20
    2: string Password (api.form="password" api.vd = "len($) > 0 && len($) < 20");//密码长度大于0小于20
}

struct LoginResp {}

struct RegisterResp {}

service userService {
    common.NilResponse Login(1: LoginReq request) (api.get="/user/login")
    common.NilResponse Register(1: RegisterReq request) (api.post="/user/register")
}