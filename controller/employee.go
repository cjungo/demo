package controller

import (
	"github.com/cjungo/cjungo"
	"github.com/labstack/echo/v4"
)

type EmployeeController struct {
}

func NewEmployeeController() *EmployeeController {
	return &EmployeeController{}
}

func (controller *EmployeeController) Detail(c echo.Context) error {
	ctx := c.(cjungo.HttpContext)
	return ctx.Resp("detail")
}
