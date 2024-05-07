package controller

import (
	"net/http"

	"github.com/cjungo/cjungo"
	"github.com/rs/zerolog"
)

type TaskController struct {
	queue  *cjungo.TaskQueue
	logger *zerolog.Logger
}

func NewTaskController(
	queue *cjungo.TaskQueue,
	logger *zerolog.Logger,
) *TaskController {
	return &TaskController{
		queue:  queue,
		logger: logger,
	}
}

type TaskPushParam struct {
	Name string         `json:"name" form:"name" query:"name"`
	Data map[string]any `json:"data" form:"data"`
}

func (controller *TaskController) Push(ctx cjungo.HttpContext) error {
	param := &TaskPushParam{}
	if err := ctx.Bind(param); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			map[string]any{
				"code":    -1,
				"message": "请求参数有误",
			},
		)
	}

	if id, err := controller.queue.PushTask(param.Name, param.Data); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			map[string]any{
				"code":    -1,
				"message": "请求参数有误",
				"id":      id,
				"error":   err,
			},
		)
	}

	return ctx.JSON(
		http.StatusOK,
		map[string]any{
			"code":    0,
			"message": "Ok",
		},
	)
}
