package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type EmployeeController struct {
}

func NewEmployeeController() *EmployeeController {
	return &EmployeeController{}
}

func (controller *EmployeeController) Detail(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusOK,
		map[string]any{
			"code":    0,
			"message": "detail",
		},
	)
}
