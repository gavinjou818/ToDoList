package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"grpc_todolist/app/gateway/internal/http"

	"github.com/gin-contrib/sessions/cookie"

	"grpc_todolist/app/gateway/middleware"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.ErrorMiddleWare())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})
		v1.POST("/user/register", http.UserRegister)
		v1.POST("/user/login", http.UserLogin)

		// 需要登录保护
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			// 任务模块
			authed.GET("task", http.GetTaskList)
			authed.POST("task", http.CreateTask)
			authed.PUT("task", http.UpdateTask)
			authed.DELETE("task", http.DeleteTask)
		}
	}
	return ginRouter
}
