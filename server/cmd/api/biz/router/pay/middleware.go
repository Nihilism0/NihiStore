// Code generated by hertz generator.

package pay

import (
	"NihiStore/server/cmd/api/config"
	"NihiStore/server/shared/middleware"
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	//return []app.HandlerFunc{
	//	middleware.JWTAuthMiddleware(config.GlobalServerConfig.JWTInfo.SigningKey),
	//}
	return nil
}

func _alipayMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _buygoodsMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.JWTAuthMiddleware(config.GlobalServerConfig.JWTInfo.SigningKey),
	}
	return nil
}

func _callbackMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _notifyMw() []app.HandlerFunc {
	// your code...
	return nil
}
