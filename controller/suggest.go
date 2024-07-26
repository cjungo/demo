package controller

import (
	"fmt"
	"time"

	"github.com/cjungo/cjungo"
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

func (controller *SuggestController) LongLongAgo(
	ctx cjungo.HttpContext,
	tx chan cjungo.LongPollingEvent,
	rx chan error,
) {
	reqId := ctx.GetReqID()
	controller.logger.Info().
		Str("action", "start").
		Str("reqId", reqId).
		Msg("[LONG POLLING DEMO]")

	for {
		select {
		case <-ctx.Request().Context().Done():
			controller.logger.Info().
				Str("action", "done").
				Str("reqId", reqId).
				Msg("[LONG POLLING DEMO]")
			return
		case err := <-rx:
			tx <- cjungo.LongPollingEvent{Err: err}
		default:
			tx <- cjungo.LongPollingEvent{Data: []byte(fmt.Sprintf("tick: %s", reqId))}
			controller.logger.Info().
				Str("action", "tick").
				Str("reqId", reqId).
				Msg("[LONG POLLING DEMO]")
			time.Sleep(4 * time.Second)
		}
	}
}

func (controller *SuggestController) Index(
	ctx cjungo.HttpContext,
	tx chan cjungo.SseEvent,
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
			tx <- cjungo.SseEvent{Data: err}
		default:
			tx <- cjungo.SseEvent{Data: fmt.Sprintf("tick: %s", reqId)}
			controller.logger.Info().
				Str("action", "tick").
				Str("reqId", reqId).
				Msg("[SSE DEMO]")
			time.Sleep(4 * time.Second)
		}
	}
}
