package controller

import (
	"time"

	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/mid"
	"github.com/cjungo/demo/misc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type LoginController struct {
}

func NewLoginController() (*LoginController, error) {
	return &LoginController{}, nil
}

func (controller *LoginController) Login(c echo.Context) error {
	ctx := c.(cjungo.HttpContext)
	claims := &misc.JwtClaims{
		UserId:   1,
		UserName: "aaaa",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}
	token, err := mid.MakeJwtToken(claims)
	if err != nil {
		return ctx.RespBad(err)
	}
	return ctx.Resp(token)
}
