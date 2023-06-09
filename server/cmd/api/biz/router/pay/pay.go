// Code generated by hertz generator. DO NOT EDIT.

package pay

import (
	pay "NihiStore/server/cmd/api/biz/handler/pay"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_alipay := root.Group("/alipay", _alipayMw()...)
		_alipay.GET("/callback", append(_callbackMw(), pay.CallBack)...)
		_alipay.POST("/notify", append(_notifyMw(), pay.Notify)...)
		_alipay.GET("/pay", append(_buygoodsMw(), pay.BuyGoods)...)
	}
}
