package misc

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	UserId   int64
	UserName string
	jwt.RegisteredClaims
}
