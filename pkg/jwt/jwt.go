package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

type Myclaims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
