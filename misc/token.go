package misc

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	EmployeeId          int32
	EmployeeNickname    string
	EmployeePermissions []int32
	jwt.RegisteredClaims
}
