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
	logger            *zerolog.Logger
	sqlite            *db.Sqlite
	captchaController *ext.CaptchaController
}

func NewLoginController(
	logger *zerolog.Logger,
	sqlite *db.Sqlite,
	captchaController *ext.CaptchaController,
) (*LoginController, error) {
	return &LoginController{
		logger:            logger,
		sqlite:            sqlite,
		captchaController: captchaController,
	}, nil
}

type LoginParam struct {
	Username      string `json:"username" form:"username" query:"username" validate:"optional" example:"admin"`
	Password      string `json:"password" form:"password" query:"password" validate:"optional" example:"admin"`
	CaptchaID     string `json:"captchaId" form:"captchaId" query:"captchaId" validate:"optional" example:"1"`
	CaptchaAnswer string `json:"captchaAnswer" form:"captchaAnswer" query:"captchaAnswer" validate:"optional" example:"12"`
}

func (controller *LoginController) Login(ctx cjungo.HttpContext) error {
	param := &LoginParam{}

	if err := ctx.Bind(param); err != nil {
		return ctx.RespBad(err)
	}

	if err := controller.captchaController.Verify(param.CaptchaID, param.CaptchaAnswer, true); err != nil {
		return ctx.RespBad(err)
	}

	employee := &model.CjEmployee{}
	password := ext.Sha256(param.Password).Hex()
	if err := controller.sqlite.Select("*").
		Where("username=? AND password=?", param.Username, password).
		Find(employee).Error; err != nil {
		return ctx.RespBad(err)
	}
	controller.logger.Info().Any("employee", employee).Str("action", "登录").Msg("[LOGIN]")

	if employee.ID == 0 {
		return ctx.RespBad("无效的账号或密码")
	}

	var permissions []string
	if err := controller.sqlite.Select("P.tag").
		Table("cj_employee_permission AS EP").
		Joins(
			"JOIN cj_permission AS P ON P.id=EP.permission_id",
		).
		Where("EP.employee_id=?", employee.ID).
		Find(&permissions).Error; err != nil {
		return ctx.RespBad(err)
	}

	controller.logger.Info().Any("permissions", permissions).Str("action", "权限").Msg("[LOGIN]")

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
