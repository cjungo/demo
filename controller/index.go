package controller

import (
	"github.com/cjungo/cjungo"
)

type IndexController struct {
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (controller *IndexController) Index(ctx cjungo.HttpContext) error {
	return ctx.RespOk()
}
