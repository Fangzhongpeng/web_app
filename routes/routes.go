package routes

import (
	"net/http"
	"web_app/dao/controller"
	"web_app/logger"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//注册路由
	r.POST("/signup", controller.SignUpHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
