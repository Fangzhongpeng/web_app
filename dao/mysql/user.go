package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"web_app/models"
)

// 把每一步数据库操作封装成函数
// 待logic层更具业务需要调用
// 不关注业务逻辑
const secret = "fangzhongpeng.com"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 向数据库中插入一条新的用户记录
//
//	func InsertUser(user *models.User) (err error) {
//		//对用户名加密
//		user.Password = GenPasswd(user.Password)
//		//执行sql入库
//		sqlStr := `insert into user(user_id, username,password) values(?,?,?)`
//		_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
//		return
//	}
//
// InsertUser 想数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行SQL语句入库
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

//	func Login(user *models.User) (err error) {
//		//oPassword := user.Password // 用户登录的密码
//		var oUser models.User
//		sqlStr := `select user_id, username, password from user where username=?`
//		err = db.Get(&oUser, sqlStr, user.Username)
//		if err == sql.ErrNoRows {
//			return ErrorUserNotExist
//		}
//		if err != nil {
//			// 查询数据库失败
//			return err
//		}
//		// 判断密码是否正确
//		// 校验密码
//		fmt.Printf("查询的用户id:%v", oUser.UserID)
//		err = ComparePasswd(oUser.Password, user.Password)
//		//err = ComparePasswd("test1111", "test")
//		if err != nil {
//			fmt.Println("错误")
//			return ErrorInvalidPassword
//		}
//		//password := encryptPassword(oPassword)
//		//if password != user.Password {
//		//	return ErrorInvalidPassword
//		//}
//		return
//	}
//
// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
func Login(user *models.User) (err error) {
	oPassword := user.Password // 用户登录的密码
	sqlStr := `select user_id, username, password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		// 查询数据库失败
		return err
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}

// GetUserById 根据id获取用户信息
func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username ,password from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	return
}

// 密码加密 使用自适应hash算法, 不可逆
//func GenPasswd(passwd string) string {
//	hashPasswd, _ := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
//	return string(hashPasswd)
//}
//func ComparePasswd(hashPasswd string, passwd string) error {
//	if err := bcrypt.CompareHashAndPassword([]byte(hashPasswd), []byte(passwd)); err != nil {
//		return err
//	}
//	return nil
//}
