package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type IndexController struct {
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (controller *IndexController) Index(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusOK,
		map[string]any{
			"code":    0,
			"message": "Ok",
		},
	)
}
