package logic

import (
	"fmt"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
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

func Login(p *models.ParamLogin) error {

	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID
	return mysql.Login(user)

}
