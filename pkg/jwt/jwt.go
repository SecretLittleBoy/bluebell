package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Myclaims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Second * 40

var mySecret = []byte("这是JWT盐")

func keyFunc(token *jwt.Token) (interface{}, error) {
	return mySecret, nil
}

func GenFullToken(userId int64, username string) (access_token, refresh_token string, err error) {
	c := Myclaims{
		UserId:   userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "bluebell",
		},
	}
	access_token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)
	if err != nil {
		return
	}
	refresh_token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(TokenExpireDuration * 24).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "bluebell",
	}).SignedString(mySecret)
	return
}

func GenAccessToken(userId int64, username string) (access_token string, err error) {
	c := Myclaims{
		UserId:   userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "bluebell",
		},
	}
	access_token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)
	return
}

func ParseAccessToken(tokenString string) (claims *Myclaims, err error) {
	var token *jwt.Token
	claims = new(Myclaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return
}

func RefreshToken(access_token, refresh_token string) (newAccessToken string, err error) {
	if _, err = jwt.Parse(refresh_token, keyFunc); err != nil {
		return
	}
	var claims Myclaims
	_, err = jwt.ParseWithClaims(access_token, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)
	if v.Errors == jwt.ValidationErrorExpired {
		newAccessToken, err = GenAccessToken(claims.UserId, claims.Username)
	}
	return
}
