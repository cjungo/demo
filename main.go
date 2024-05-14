package main

import (
	"log"
	"net/http"
	"time"

	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/cjungo/ext"
	"github.com/cjungo/cjungo/mid"
	"github.com/cjungo/demo/controller"
	"github.com/cjungo/demo/entity"
	localEntity "github.com/cjungo/demo/local/entity"
	localModel "github.com/cjungo/demo/local/model"
	"github.com/cjungo/demo/misc"
	"github.com/cjungo/demo/model"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	_ "github.com/cjungo/demo/docs"
)

func route(
	router cjungo.HttpRouter,
	logger *zerolog.Logger,
	permitManager *mid.PermitManager[int32, misc.EmployeeToken],
	indexController *controller.IndexController,
	loginController *controller.LoginController,
	taskController *controller.TaskController,
	employeeController *controller.EmployeeController,
	productController *controller.ProductController,
) http.Handler {
	router.GET("/", indexController.Index)
	router.POST("/login", loginController.Login)

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
	productGroup.PUT("/add", productController.Add, permitManager.Permit(1))
	productGroup.GET("/detail", productController.Detail, permitManager.Permit(1, 2, 4))
	productGroup.POST("/edit", productController.Edit, permitManager.Permit(1))

	return router.GetHandler()
}

func init() {
	if err := cjungo.LoadEnv(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	app, err := cjungo.NewApplication(func(c cjungo.DiContainer) error {
		// 加载日志配置
		if err := c.Provide(cjungo.LoadLoggerConfFromEnv); err != nil {
			return err
		}

		// 加载数据库配置
		if err := c.Provide(db.LoadMySqlConfFormEnv); err != nil {
			return err
		}
		if err := c.Provide(db.LoadSqliteConfFormEnv); err != nil {
			return err
		}

		// 加载服务器配置
		if err := c.Provide(cjungo.LoadHttpServerConfFromEnv); err != nil {
			return err
		}

		// 加载队列配置
		if err := c.Provide(cjungo.LoadTaskConfFromEnv); err != nil {
			return err
		}

		// 注册权限管理器
		if err := c.Provide(mid.NewPermitManager(func(ctx cjungo.HttpContext) ([]int32, misc.EmployeeToken, error) {
			claims := &misc.JwtClaims{}
			if _, err := ext.ParseJwtToken(ctx, claims); err != nil {
				return nil, claims.EmployeeToken, err
			}
			return claims.EmployeePermissions, claims.EmployeeToken, nil
		})); err != nil {
			return err
		}

		// 注册数据库
		if err := c.Provide(db.NewMySqlHandle(func(mysql *db.MySql) error {
			entity.Use(mysql.DB)
			mysql.AutoMigrate(&model.CjProduct{})
			return nil
		})); err != nil {
			return err
		}
		if err := c.Provide(db.NewSqliteHandle(func(sqlite *db.Sqlite) error {
			localEntity.Use(sqlite.DB)
			sqlite.AutoMigrate(
				&localModel.CjPermission{},
				&localModel.CjOperation{},
				&localModel.CjEmployee{},
				&localModel.CjEmployeePermission{},
			)
			return sqlite.Transaction(func(tx *gorm.DB) error {
				now := time.Now()
				admin := &localModel.CjEmployee{
					ID:       1,
					Username: "admin",
					Password: "admin",
					Nickname: "admin",
					CreateBy: 0,
					CreateAt: now,
					UpdateBy: 0,
					UpdateAt: now,
				}
				tx.Save(admin)

				return nil
			})
		})); err != nil {
			return err
		}

		// 注册队列
		if err := c.Provide(cjungo.NewTaskQueueHandle(func(queue *cjungo.TaskQueue) error {
			queue.RegisterProcess("action-1", func(param *cjungo.TaskAction) (cjungo.TaskResultMessage, error) {
				queue.Logger.Info().
					Str("name", param.Name).
					Str("id", param.ID).
					Msg("任务执行完成")
				return nil, nil
			})
			return nil
		})); err != nil {
			return err
		}

		// 注册控制器
		if err := c.ProvideController([]any{
			controller.NewIndexController,
			controller.NewLoginController,
			controller.NewTaskController,
			controller.NewEmployeeController,
			controller.NewProductController,
		}); err != nil {
			return err
		}

		// 注册路由
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
