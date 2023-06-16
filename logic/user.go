package logic

import (
	"fmt"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"

	"go.uber.org/zap"
)

// 存放业务逻辑代码
func SignUp(p *models.ParamSignUp) (err error) {
	//1 判断用户存不存在

	if err := mysql.CheckUserExist(p.Username); err != nil {
		//数据库查询出错
		//fmt.Println("查询用户出错")
		return err
	}
	//2 生成UID
	userID := snowflake.GenID()
	// 构造一个user实例
	fmt.Printf("%s", userID)
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//3 保存到数据库

	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (token string, err error) {

	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	//fmt.Printf("用户id是：%v", user.UserID)
	//生成jwt token
	return jwt.GenToken(user.UserID, user.Username)
}
func GetUserById(uid int64) (data *models.User, err error) {
	// 查询并组合我们接口想用的数据
	// 根据作者id查询作者信息
	user, err := mysql.GetUserById(uid)
	if err != nil {
		//ResponseError(c, CodeNeedLogin)
		zap.L().Error("logic.Login failed", zap.Int64("userid", user.UserID), zap.Error(err))
		return
	}

	// 接口数据拼接
	data = &models.User{
		UserID:   user.UserID,
		Username: user.Username,
		Password: user.Password,
	}
	return
}
