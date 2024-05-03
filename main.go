package main

import (
	"log"
	"net/http"

	"github.com/cjungo/cjungo"
	"github.com/cjungo/demo/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/dig"
)

func provideController(container *dig.Container) error {
	controllers := []any{
		controller.NewIndexController,
		controller.NewEmployeeController,
	}
	for _, c := range controllers {
		if err := container.Provide(c); err != nil {
			return err
		}
	}
	return nil
}

func route(
	e *echo.Echo,
	indexController *controller.IndexController,
	employeeController *controller.EmployeeController,
) http.Handler {
	e.GET("/", indexController.Index)
	employeeGroup := e.Group("employee", middleware.Gzip())
	employeeGroup.GET("/detail", employeeController.Detail)

	return e
}

func main() {
	app, err := cjungo.NewApplication(func(c *dig.Container) error {
		// 加载日志配置
		if err := c.Provide(cjungo.LoadLoggerConfFromEnv); err != nil {
			return err
		}
		// 加载服务器配置
		if err := c.Provide(cjungo.LoadHttpServerConfFromEnv); err != nil {
			return err
		}
		if err := provideController(c); err != nil {
			return err
		}
		if err := c.Provide(route); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
