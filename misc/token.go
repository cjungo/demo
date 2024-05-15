package misc

import (
	"github.com/golang-jwt/jwt/v5"
)

type EmployeeToken struct {
	EmployeeId          int32    `json:"eid,omitempty"`
	EmployeeNickname    string   `json:"nickname,omitempty"`
	EmployeePermissions []string `json:"permissions,omitempty"`
}

type JwtClaims struct {
	EmployeeToken
	jwt.RegisteredClaims
}

func (claims *JwtClaims) GetPermissions() []string {
	return claims.EmployeePermissions
}

func (claims *JwtClaims) GetToken() EmployeeToken {
	return claims.EmployeeToken
}
