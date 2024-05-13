package controller

import (
	"time"

	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/cjungo/ext"
	"github.com/cjungo/demo/misc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog"

	"github.com/cjungo/demo/local/model"
)

type LoginController struct {
	logger *zerolog.Logger
	sqlite *db.Sqlite
}

func NewLoginController(
	logger *zerolog.Logger,
	sqlite *db.Sqlite,
) (*LoginController, error) {
	return &LoginController{
		logger: logger,
		sqlite: sqlite,
	}, nil
}

type LoginParam struct {
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}

func (controller *LoginController) Login(ctx cjungo.HttpContext) error {
	param := &LoginParam{}

	if err := ctx.Bind(param); err != nil {
		return ctx.RespBad(err)
	}

	employee := &model.CjEmployee{}
	if err := controller.sqlite.Select("*").
		Where("username=? AND password=?", param.Username, param.Password).
		Find(employee).Error; err != nil {
		return ctx.RespBad(err)
	}
	controller.logger.Info().Any("employee", employee).Msg("登录")

	if employee.ID == 0 {
		return ctx.RespBad("无效的账号或密码")
	}

	var permissions []int32
	if err := controller.sqlite.Select("permission_id").
		Table("cj_employee_permission").
		Where("employee_id=?", employee.ID).
		Find(&permissions).Error; err != nil {
		return ctx.RespBad(err)
	}

	controller.logger.Info().Any("permissions", permissions).Msg("权限")

	claims := &misc.JwtClaims{
		EmployeeToken: misc.EmployeeToken{
			EmployeeId:          employee.ID,
			EmployeeNickname:    employee.Nickname,
			EmployeePermissions: permissions,
		},
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
	token, err := ext.MakeJwtToken(claims)
	if err != nil {
		return ctx.RespBad(err)
	}
	return ctx.Resp(token)
}
