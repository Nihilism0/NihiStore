package middleware

import (
	"NihiStore/server/shared/model"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

func JWTAuthMiddleware(Secret string) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, utils.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, utils.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(parts[1], Secret)
		if err != nil {
			c.JSON(http.StatusOK, utils.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		c.Set("ID", mc.Id)
		c.Next(ctx)
	}
}

func ParseToken(tokenString, Secret string) (*model.CustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

//func NewClaim(id uint) model.CustomClaims {
//	claim := model.CustomClaims{
//		ID: id, // 自定义字段
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // 过期时间
//			Issuer:    consts.JWTIssuer,                          // 签发人
//		},
//	}
//	return claim
//}

func NewJWT(signingKey string) *JWT {
	return &JWT{
		SigningKey: []byte(signingKey),
	}
}

type JWT struct {
	SigningKey []byte
}
