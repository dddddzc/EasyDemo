package common

import (
	"easydemo/model"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("easydemo_secret_key")

// Claims 定义一个结构体，用来存储要加密的信息
type Claims struct {
	// 此处的ID是GORM MODEL中的ID
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	// 过期时间点
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	// 生成token
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),     // 请求时间
			Issuer:    "easydemo",
			Subject:   "user token",
		},
	}
	// 使用256位的密钥签名token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
