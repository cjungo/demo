package controller

import (
	"github.com/cjungo/cjungo"
)

type IndexController struct {
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

// Index godoc
// @Summary      index
// @Description  home
// @Tags         index
// @Produce      json
// @Success      200  {object}  any
// @Failure      400  {object}  error
// @Router       / [get]
func (controller *IndexController) Index(ctx cjungo.HttpContext) error {
	return ctx.RespOk()
}

func (controller *IndexController) Status(ctx cjungo.HttpContext) error {
	return ctx.Resp(map[string]any{
		"reqID":  ctx.GetReqID(),
		"reqAt":  ctx.GetReqAt(),
		"realIP": ctx.RealIP(),
	})
}
