package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// TODO 基本不理解
var mySigningKey = []byte("signKey")

func GenJWT() (string, error) {
	claims := &jwt.RegisteredClaims{
		Issuer:    "bluebell_practice",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 3650)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func ParseJWT(tokenString string) bool {
	// 解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil { // 解析token失败
		return false
	}
	return token.Valid
}
