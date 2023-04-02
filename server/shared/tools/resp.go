package tools

import (
	"NihiStore/server/shared/kitex_gen/base"
)

func BuildBaseResp(code int64, msg string) *base.BaseResponse {
	return &base.BaseResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}
}

func BuildToken() {

}
