package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/mid"
	"github.com/cjungo/demo/controller"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"go.uber.org/dig"
)

func provideController(container *dig.Container) error {
	controllers := []any{
		controller.NewIndexController,
		controller.NewLoginController,
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
	logger *zerolog.Logger,
	indexController *controller.IndexController,
	loginController *controller.LoginController,
	employeeController *controller.EmployeeController,
) http.Handler {
	e.GET("/", indexController.Index)
	e.POST("/login", loginController.Login)

	// api 加了 JWT 权限验证
	apiGroup := e.Group("/api", mid.NewJwtAuthMiddleware(func(token *jwt.Token) error {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// 根据业务需求，验证凭证里面的信息
			logger.Info().Str("claims", fmt.Sprintf("%v", claims)).Msg("claims:")
		} else {
			return fmt.Errorf("获取凭证失败: %v", token.Claims)
		}
		return nil
	}))
	employeeGroup := apiGroup.Group("/employee", middleware.Gzip())
	employeeGroup.GET("/detail", employeeController.Detail)

	return e
}

func init() {
	if err := cjungo.LoadEnv(); err != nil {
		log.Fatalln(err)
	}
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
