package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MyClaims struct {
	UserId int   `json:"userId"`
	ExpAt  int64 `json:"expAt"`
	jwt.RegisteredClaims
}

// GetJwtToken 获取jwt token
func GetJwtToken(secret string, seconds int64, userId int) (string, error) {
	claim := MyClaims{
		UserId: userId,
		ExpAt:  time.Now().Add(time.Second * time.Duration(seconds)).Unix(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(seconds))), // 过期时间7天
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                           // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                           // 生效时间
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析token
func ParseToken(tokenString, secret string) (*MyClaims, error) {
	t, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := t.Claims.(*MyClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
