package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func GenerateToken(username string, userId int64) (string, error) {
	now_time := time.Now()
	expireTime := now_time.Add(time.Hour)

	claims := Claims{
		userId,
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("golang"))
	if err != nil {
		return "err", err
	}
	return token, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte("golang"), nil
		})
	if err != nil {
		return nil, nil, err
	}
	claim := token.Claims.(*Claims)
	return token, claim, err
}
