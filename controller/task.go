package controller

import (
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
		return ctx.RespBad(err)
	}

	if id, err := controller.queue.PushTask(param.Name, param.Data); err != nil {
		return ctx.RespBadF("任务分发失败 ID: %s %v", id, err)
	} else {
		return ctx.Resp(id)
	}
}

func (controller *TaskController) Query(ctx cjungo.HttpContext) error {
	id := ctx.QueryParam("id")
	controller.logger.Info().
		Str("ID", id).
		Msg("任务")
	if result, err := controller.queue.QueryTask(id); err != nil {
		return ctx.RespBad(err)
	} else {
		return ctx.Resp(result)
	}
}

