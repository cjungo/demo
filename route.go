package main

// 路由示例

import (
	"net/http"

	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/ext"
	"github.com/cjungo/cjungo/mid"
	"github.com/cjungo/demo/controller"
	"github.com/cjungo/demo/misc"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

// 注册控制器类列表
var provideControllers = []any{
	controller.NewIndexController,
	controller.NewLoginController,
	controller.NewTaskController,
	controller.NewEmployeeController,
	controller.NewProductController,
	ext.NewCaptchaController,
}

// 路由注册
func route(
	router cjungo.HttpRouter,
	logger *zerolog.Logger,
	permitManager *mid.PermitManager[string, misc.EmployeeToken],
	captchaController *ext.CaptchaController,
	indexController *controller.IndexController,
	loginController *controller.LoginController,
	taskController *controller.TaskController,
	employeeController *controller.EmployeeController,
	productController *controller.ProductController,
) http.Handler {
	router.GET("/", indexController.Index)
	router.GET("/status", indexController.Status)
	router.POST("/login", loginController.Login)
	router.GET("/captcha/math", captchaController.GenerateMath)

	apiGroup := router.Group("/api")

	// task
	taskGroup := apiGroup.Group("/task", middleware.CORS())
	taskGroup.POST("/push", taskController.Push)
	taskGroup.GET("/query", taskController.Query)

	// employee
	employeeGroup := apiGroup.Group("/employee", middleware.Gzip())
	employeeGroup.PUT("/add", employeeController.Add)
	employeeGroup.GET("/detail", employeeController.Detail)
	employeeGroup.POST("/edit", employeeController.Edit)

	// product
	productGroup := apiGroup.Group("/product")
	productGroup.PUT("/add", productController.Add, permitManager.Permit("product_add"))
	productGroup.GET("/detail", productController.Detail, permitManager.Permit("product_find"))
	productGroup.POST("/edit", productController.Edit, permitManager.Permit("product_edit"))

	return router.GetHandler()
}
