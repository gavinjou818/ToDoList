package middleware

import (
	"github.com/gin-gonic/gin"
	"grpc_todolist/pkg/ctl"
	"grpc_todolist/pkg/e"
	"grpc_todolist/pkg/util/jwt"
	"time"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})
			c.Abort()
		}
		claims, err := jwt.ParseToken(token) // 要从请求中带过来
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}

		// 相当于说，通过 ctx，把用户信息放入，获取新的 context
		c.Request = c.Request.WithContext(ctl.NewContext(c.Request.Context(), &ctl.UserInfo{Id: claims.UserID}))
		ctl.InitUserInfo(c.Request.Context())
		c.Next()
	}
}
