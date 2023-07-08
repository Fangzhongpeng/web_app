package router

import (
	"net/http"
	"web_app/controller"
	"web_app/logger"
	"web_app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	// 启用全局跨域中间件
	r.Use(middleware.CORSMiddleware())
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//注册路由
	v1 := r.Group("/api/v1")

	// 注册
	//v1.Use(middleware.CORSMiddleware())
	v1.POST("/signup", controller.SignUpHandler)
	// 登录
	v1.POST("/login", controller.LoginHandler)
	//v1.Use(middleware.JWTAuthMiddleware())

	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.GET("/userinfo", controller.GetUserInfoHandler)
		v1.GET("/community", controller.CommunityHandler)
		v1.POST("posts", controller.CreatePostHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostListHandler)
		v1.GET("/posts2", controller.GetPostListHandler2)
		v1.POST("/vote", controller.PostVoteController)
		v1.GET("/ping", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
			//如果是登录的用户,判断请求头中是否有有效的jwt token
			c.String(http.StatusOK, "pong")
		})
	}

	//r.Use(middleware.JWTAuthMiddleware())
	//{
	//	r.GET("/userinfo", controller.GetUserInfoHandler)
	//}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
