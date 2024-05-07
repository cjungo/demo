package controller

import (
	"github.com/cjungo/cjungo"
	"github.com/labstack/echo/v4"
)

type IndexController struct {
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (controller *IndexController) Index(c echo.Context) error {
	ctx := c.(cjungo.HttpContext)
	return ctx.RespOk()
}
