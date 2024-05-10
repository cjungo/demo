package misc

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	EmployeeId          int32   `json:"eid,omitempty"`
	EmployeeNickname    string  `json:"nickname,omitempty"`
	EmployeePermissions []int32 `json:"permissions,omitempty"`
	jwt.RegisteredClaims
}
