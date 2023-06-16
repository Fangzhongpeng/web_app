package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	// 可根据需要自行添加字段
	UserID               int64  `json:"user_id"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

//type MyClaims struct {
//	UserID   int64  `json:"user_id"`
//	Username string `json:"username"`
//	jwt.StandardClaims
//}

const TokenExpireDuration = time.Hour * 2

var CustomSecret = []byte("G小调交响曲")

//GenToken 生成JWT

func GenToken(userID int64, username string) (string, error) {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		userID,
		username, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "web_app", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(CustomSecret)
}

//func GenToken(userID int64, username string) (string, error) {
//	// 创建一个我们自己的声明的数据
//	c := MyClaims{
//		userID,
//		"username", // 自定义字段
//		jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(
//				time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(), // 过期时间
//			Issuer: "bluebell", // 签发人
//		},
//	}
//	// 使用指定的签名方法创建签名对象
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
//	// 使用指定的secret签名并获得完整的编码后的字符串token
//	return token.SignedString(CustomSecret)
//}

//// ParseToken 解析JWT
//func ParseToken(tokenString string) (*CustomClaims, error) {
//	// 解析token
//	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
//	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
//		// 直接使用标准的Claim则可以直接使用Parse方法
//		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
//		return CustomSecret, nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	// 对token对象中的Claim进行类型断言
//	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
//		return claims, nil
//	}
//	return nil, errors.New("invalid token")
//}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	var mc = new(CustomClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
