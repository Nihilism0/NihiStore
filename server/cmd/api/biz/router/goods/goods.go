// Code generated by hertz generator. DO NOT EDIT.

package goods

import (
	goods "NihiStore/server/cmd/api/biz/handler/goods"
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
		_goods := root.Group("/goods", _goodsMw()...)
		_goods.POST("/create", append(_creategoodsMw(), goods.CreateGoods)...)
		_goods.DELETE("/delete", append(_deletegoodsMw(), goods.DeleteGoods)...)
		_goods.GET("/getgoodsphoto", append(_getgoodsphotoMw(), goods.GetGoodsPhoto)...)
		_goods.GET("/searchgoods", append(_searchgoodsMw(), goods.SearchGoods)...)
		_goods.GET("/searchgoodsinfo", append(_searchgoodsinfoMw(), goods.SearchGoodsInfo)...)
		_goods.PUT("/update", append(_updategoodsMw(), goods.UpdateGoods)...)
		_goods.GET("/uploadgoodsphoto", append(_uploadgoodsphotoMw(), goods.UploadGoodsPhoto)...)
	}
}
