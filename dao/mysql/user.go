package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"web_app/models"

	"golang.org/x/crypto/bcrypt"
)

// 把每一步数据库操作封装成函数
// 待logic层更具业务需要调用
// 不关注业务逻辑

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	//对用户名加密
	user.Password = GenPasswd(user.Password)
	//执行sql入库
	sqlStr := `insert into user(user_id, username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func Login(user *models.User) (err error) {
	//oPassword := user.Password // 用户登录的密码
	var oUser models.User
	sqlStr := `select user_id, username, password from user where username=?`
	err = db.Get(&oUser, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return err
	}
	if err != nil {
		// 查询数据库失败
		return err
	}
	// 判断密码是否正确
	// 校验密码
	err = ComparePasswd(oUser.Password, user.Password)
	//err = ComparePasswd("test1111", "test")
	if err != nil {
		fmt.Println("错误")
		return errors.New("密码错误")
	}
	//password := encryptPassword(oPassword)
	//if password != user.Password {
	//	return ErrorInvalidPassword
	//}
	return
}

// 密码加密 使用自适应hash算法, 不可逆
func GenPasswd(passwd string) string {
	hashPasswd, _ := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	return string(hashPasswd)
}
func ComparePasswd(hashPasswd string, passwd string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPasswd), []byte(passwd)); err != nil {
		return err
	}
	return nil
}
