package controller

import (
	"errors"
	"fmt"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": err.Error(),
			//})
			ResponseError(c, CodeInvalidParam)
			return
		}
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		//})
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	fmt.Println(p)
	// 2，业务处理
	if err := logic.SignUp(p); err != nil {
		//fmt.Printf("%s", err)
		zap.L().Error("logic.SignUp faild", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": "注册失败",
		//})
		return
	}

	// 3，返回响应

	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "success",
	//})
	ResponseSuccess(c, nil)
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
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": err.Error(),
			//})
			ResponseError(c, CodeInvalidParam)
			return
		}
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		//})
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2 业务逻辑处理
	user, err := logic.Login(p)

	if err != nil {
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": "用户名或密码错误",
		//})
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	//3 返回响应
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "登录成功",
	//})
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID), // id值大于1<<53-1  int64类型的最大值是1<<63-1
		"user_name": user.Username,
		"token":     user.Token,
	})
}
func GetUserInfoHandler(c *gin.Context) {
	userID, err := getCurrentUserID(c)
	//fmt.Printf("获取的用户id是啊:%v", userID)

	if err != nil {
		//ResponseError(c, CodeNeedLogin)
		zap.L().Error("logic.Login failed", zap.String("username", "p.Username"), zap.Error(err))
		return
	}
	data, err := logic.GetUserById(userID)
	ResponseSuccess(c, data)
}