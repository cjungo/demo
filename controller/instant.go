package controller

import (
	"fmt"
	"sync"

	"github.com/cjungo/cjungo"
	"github.com/rs/zerolog"
	"golang.org/x/net/websocket"
)

type InstantController struct {
	connections sync.Map
	logger      *zerolog.Logger
}

func NewInstantController(logger *zerolog.Logger) *InstantController {
	return &InstantController{
		logger: logger,
	}
}

func (controller *InstantController) sendTo(id string, msg string) error {
	if v, ok := controller.connections.Load(id); ok {
		conn := v.(*websocket.Conn)
		return websocket.Message.Send(conn, msg)
	}
	return fmt.Errorf("无效ID: %s", id)
}

func (controller *InstantController) Index(ctx cjungo.HttpContext) error {
	websocket.Handler(func(conn *websocket.Conn) {
		id := ctx.GetReqID()
		controller.connections.Store(id, conn)

		defer func() {
			conn.Close()
		}()

		go func() {
			for {
				msg := ""
				if err := websocket.Message.Receive(conn, &msg); err != nil {
					controller.sendTo(id, msg)
				} else {
					controller.logger.Error().Err(err).Msg("[INSTANT]")
				}
			}
		}()
	}).ServeHTTP(ctx.Response(), ctx.Request())
	return nil
}
