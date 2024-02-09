package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"errors"
)

type Myclaims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var mySecret = []byte("这是JWT盐")

func GenToken(userId int64, username string) (string, error) {
	c := Myclaims{
		UserId:   userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "bluebell",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(mySecret)
}

func ParseToken(tokenString string) (*Myclaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Myclaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Myclaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
