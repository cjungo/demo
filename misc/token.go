package misc

import (
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/mid"
	"github.com/golang-jwt/jwt/v5"
)

// 管理 TOKEN 封装示例，可根据业务，存储信息，提供共用的方法。
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

func (claims *JwtClaims) GetStore() EmployeeToken {
	return claims.EmployeeToken
}

type JwtClaimsManager struct {
	permitManager *mid.PermitManager[string, EmployeeToken]
}

func NewJwtClaimsManger(
	permitManager *mid.PermitManager[string, EmployeeToken],
) *JwtClaimsManager {
	return &JwtClaimsManager{
		permitManager: permitManager,
	}
}

func (manager *JwtClaimsManager) GetToken(ctx cjungo.HttpContext) EmployeeToken {
	proof, ok := manager.permitManager.GetProof(ctx)
	if !ok {
		panic("无效TOKEN ID")
	}
	return proof.GetStore()
}
