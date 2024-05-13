package misc

import (
	"github.com/golang-jwt/jwt/v5"
)

type EmployeeToken struct {
	EmployeeId          int32   `json:"eid,omitempty"`
	EmployeeNickname    string  `json:"nickname,omitempty"`
	EmployeePermissions []int32 `json:"permissions,omitempty"`
}

type JwtClaims struct {
	EmployeeToken
	jwt.RegisteredClaims
}
