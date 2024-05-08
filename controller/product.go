package controller

import (
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"go.uber.org/dig"
)

type ProductController struct {
	mysql *db.MySql
}

type ProductControllerDi struct {
	dig.In
	MySql *db.MySql `optional:"true"`
}

func NewProductController(di ProductControllerDi) (*ProductController, error) {
	return &ProductController{
		mysql: di.MySql,
	}, nil
}

func (controller *ProductController) Detail(ctx cjungo.HttpContext) error {
	if controller.mysql != nil {
		return ctx.RespOk()
	} else {
		return ctx.RespOk()
	}
}
