package controller

import (
	"fmt"
	"net/http"
	"web_app/models"

	"go.uber.org/zap"

	"web_app/logic"

	"github.com/gin-gonic/gin"
)

// 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1,获取参数，参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Signup with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	fmt.Println(p)
	// 2，业务处理
	// 3，返回响应
	logic.SignUp()
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
