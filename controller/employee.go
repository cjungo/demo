package controller

import (
	"github.com/cjungo/cjungo"
)

type EmployeeController struct {
}

func NewEmployeeController() *EmployeeController {
	return &EmployeeController{}
}

func (controller *EmployeeController) Detail(ctx cjungo.HttpContext) error {
	return ctx.Resp("detail")
}
