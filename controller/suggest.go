package controller

import (
	"fmt"
	"time"

	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/ext"
	"github.com/rs/zerolog"
)

type SuggestController struct {
	logger *zerolog.Logger
}

func NewSuggestController(
	logger *zerolog.Logger,
) *SuggestController {
	return &SuggestController{
		logger: logger,
	}
}

func (controller *SuggestController) SseDispatch(
	ctx cjungo.HttpContext,
	tx chan ext.SseEvent,
	rx chan error,
) {
	reqId := ctx.GetReqID()
	controller.logger.Info().
		Str("action", "start").
		Str("reqId", reqId).
		Msg("[SSE DEMO]")
	for {
		select {
		case <-ctx.Request().Context().Done():
			controller.logger.Info().
				Str("action", "done").
				Str("reqId", reqId).
				Msg("[SSE DEMO]")
			return
		case err := <-rx:
			tx <- ext.SseEvent{Data: err}
		default:
			tx <- ext.SseEvent{Data: fmt.Sprintf("tick: %s", reqId)}
			controller.logger.Info().
				Str("action", "tick").
				Str("reqId", reqId).
				Msg("[SSE DEMO]")
			time.Sleep(4 * time.Second)
		}
	}
}
