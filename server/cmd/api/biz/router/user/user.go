// Code generated by hertz generator. DO NOT EDIT.

package user

import (
	user "NihiStore/server/cmd/api/biz/handler/user"
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
		_user := root.Group("/user", _userMw()...)
		_user.GET("/login", append(_loginMw(), user.Login)...)
		_user.POST("/request", append(_registerMw(), user.Register)...)
	}
}
