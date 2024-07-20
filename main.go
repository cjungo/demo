package main

import (
	"log"

	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/cjungo/ext"
	"github.com/cjungo/cjungo/mid"
	"github.com/cjungo/demo/entity"
	localEntity "github.com/cjungo/demo/local/entity"
	localModel "github.com/cjungo/demo/local/model"
	"github.com/cjungo/demo/misc"
	"github.com/cjungo/demo/model"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	_ "github.com/cjungo/demo/docs"
)

func init() {
	if err := cjungo.LoadEnv(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	app, err := cjungo.NewApplication(func(c cjungo.DiContainer) error {
		if err := c.Provides(
			cjungo.LoadLoggerConfFromEnv,      // 加载日志配置
			db.LoadMySqlConfFormEnv,           // 加载 Mysql 配置
			db.LoadSqliteConfFormEnv,          // 加载 Sqlite 配置
			cjungo.LoadHttpServerConfFromEnv,  // 加载服务器配置
			cjungo.LoadTaskConfFromEnv,        // 加载队列配置
			ext.NewStorageManager,             // 加载存储管理
			misc.ProvideMyMessageController(), // 加载 MyMessage 控制器
		); err != nil {
			return err
		}

		// 执行一些操作。
		if err := c.Invoke(func(logger *zerolog.Logger) {
			logger.Info().
				Str("action", "math").
				Any("max", cjungo.Max(123, 456)).
				Any("maxOf", cjungo.MaxOf(123, 4, 556, 767, 888, 222, 33, 1)).
				Any("min", cjungo.Min(123, 345)).
				Any("minOf", cjungo.MinOf(123, 445, 6532, 1, 343, -1)).
				Any("limit", cjungo.Limit(1, 100, 200)).
				Any("limitOf", cjungo.LimitOf(1, 100, 12, 235, 235, 47, 0, 6, 98, 23)).
				Msg("[UTIL]")
		}); err != nil {
			return err
		}

		// 注册权限管理器，依赖 misc.NewJwtClaimsManger
		if err := c.Provide(mid.NewPermitManager(func(ctx cjungo.HttpContext) (mid.PermitProof[string, misc.EmployeeToken], error) {
			claims := &misc.JwtClaims{}
			if _, err := ext.ParseJwtToken(ctx, claims); err != nil {
				return claims, &cjungo.ApiError{
					Code:     -4,
					Message:  "TOKEN 无效",
					HttpCode: 400,
					Reason:   err,
				}
			}
			return claims, nil
		})); err != nil {
			return err
		}
		if err := c.Provide(misc.NewJwtClaimsManger); err != nil {
			return err
		}

		// 注册数据库，依赖 db.LoadMySqlConfFormEnv
		if err := c.Provide(db.NewMySqlHandle(func(mysql *db.MySql) error {
			entity.Use(mysql.DB)
			mysql.AutoMigrate(&model.CjProduct{})
			return nil
		})); err != nil {
			return err
		}
		// 依赖 db.LoadSqliteConfFormEnv
		if err := c.Provide(db.NewSqliteHandle(func(sqlite *db.Sqlite) error {
			localEntity.Use(sqlite.DB)
			sqlite.AutoMigrate(
				&localModel.CjPermission{},
				&localModel.CjOperation{},
				&localModel.CjEmployee{},
				&localModel.CjEmployeePermission{},
			)
			return sqlite.Transaction(func(tx *gorm.DB) error {
				if err := misc.EnsurePermissions(tx); err != nil {
					return err
				}

				if err := misc.EnsureAdmin(tx); err != nil {
					return err
				}

				return nil
			})
		})); err != nil {
			return err
		}

		// 注册 ETCD 发现服务
		if err := c.Provide(ext.NewEtcdDiscovery); err != nil {
			return err
		}
		// 注册 ETCD 注册服务
		if err := c.Provide(ext.NewEtcdRegister); err != nil {
			return nil
		}

		// 注册队列，依赖 cjungo.LoadTaskConfFromEnv
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
		if err := c.Provides(provideControllers...); err != nil {
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
	app.BeforeRun = func(container cjungo.DiContainer) error {
		// ETCD 注册服务
		if err := registerEtcdService(container); err != nil {
			return err
		}

		// ETCD 发现服务
		return registerEtcdDiscovery(container)
	}
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
