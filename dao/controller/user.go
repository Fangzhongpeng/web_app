package controller

import (
	"fmt"
	"net/http"
	"web_app/logic"
	"web_app/models"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1,获取参数，参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Signup with invalid param", zap.Error(err))
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		})
		return
	}
	fmt.Println(p)
	// 2，业务处理
	if err := logic.SignUp(p); err != nil {
		//fmt.Printf("%s", err)
		zap.L().Error("logic.SignUp faild", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}

	// 3，返回响应

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
func LoginHandler(c *gin.Context) {
	// 1 获取请求参数以及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		})
		return
	}
	// 2 业务逻辑处理
	err := logic.Login(p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码错误",
		})
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		return
	}
	// 3 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}
