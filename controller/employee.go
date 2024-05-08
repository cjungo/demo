package controller

import (
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
)

type EmployeeController struct {
	sqlite *db.Sqlite
}

func NewEmployeeController(
	sqlite *db.Sqlite,
) *EmployeeController {
	return &EmployeeController{
		sqlite: sqlite,
	}
}

func (controller *EmployeeController) Detail(ctx cjungo.HttpContext) error {

	return ctx.Resp("detail")
}
