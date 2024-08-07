package main

// 路由示例

import (
	"net/http"
	"os"
	"path/filepath"

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
	controller.NewSuggestController,
	controller.NewInstantController,
}

// 路由注册
func route(
	router cjungo.HttpRouter,
	logger *zerolog.Logger,
	storageManager *ext.StorageManager,
	permitManager *mid.PermitManager[string, misc.EmployeeToken],
	captchaController *ext.CaptchaController,
	indexController *controller.IndexController,
	loginController *controller.LoginController,
	taskController *controller.TaskController,
	employeeController *controller.EmployeeController,
	productController *controller.ProductController,
	instantController *controller.InstantController,
	messageController *misc.MyMessageController,
	suggestController *controller.SuggestController,
) (http.Handler, error) {
	here, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOriginFunc: func(origin string) (bool, error) {
			return true, nil
		},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	router.GET("/", indexController.Index)
	router.GET("/status", indexController.Status)
	router.POST("/login", loginController.Login)
	router.GET("/captcha/math", captchaController.GenerateMath)

	uploadDir := filepath.Join(here, "upload")
	storageManager.Route(router, &ext.StorageConf{
		PathPrefix: "/upload",
		Dir:        uploadDir,
	})

	apiGroup := router.Group("/api")

	// 推荐
	suggestGroup := apiGroup.Group("/suggest")
	suggestGroup.SSE("/index", suggestController.Index)
	suggestGroup.LongPolling("/ago", suggestController.LongLongAgo)

	// 消息
	router.GET("/msg", messageController.Dispatch)

	// instant
	apiGroup.GET("/instant", instantController.Index)

	// task
	taskGroup := apiGroup.Group("/task", middleware.CORS())
	taskGroup.POST("/push", taskController.Push)
	taskGroup.GET("/query", taskController.Query)

	// employee
	employeeGroup := apiGroup.Group("/employee", middleware.Gzip())
	employeeGroup.PUT("/add", employeeController.Add)
	employeeGroup.GET("/detail", employeeController.Detail)
	employeeGroup.POST("/edit", employeeController.Edit)
	employeeGroup.DELETE("/delete", employeeController.Delete)

	// product
	productGroup := apiGroup.Group("/product")
	productGroup.PUT("/add", productController.Add, permitManager.Permit("product_add"), permitManager.Permit("product_find")) // AND
	productGroup.GET("/detail", productController.Detail, permitManager.Permit("product_find"))
	productGroup.POST("/edit", productController.Edit, permitManager.Permit("product_add", "product_edit")) // OR

	return router.GetHandler(), nil
}
