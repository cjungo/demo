package main

import (
	"log"
	"time"

	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/cjungo/ext"
	"github.com/cjungo/cjungo/mid"
	"github.com/cjungo/demo/entity"
	localEntity "github.com/cjungo/demo/local/entity"
	localModel "github.com/cjungo/demo/local/model"
	"github.com/cjungo/demo/misc"
	"github.com/cjungo/demo/model"
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
		if err := c.Provide(mid.NewPermitManager(func(ctx cjungo.HttpContext) (mid.PermitProof[string, misc.EmployeeToken], error) {
			claims := &misc.JwtClaims{}
			if _, err := ext.ParseJwtToken(ctx, claims); err != nil {
				return claims, err
			}
			return claims, nil
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
		if err := c.ProvideController(provideControllers); err != nil {
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
		if err := container.Invoke(func(register *ext.EtcdRegister) error {
			pairs := []ext.EtcdPair{
				{Key: "/web/kkkkk1", Value: "一些信息2"},
				{Key: "/web/kkkkk2", Value: "一些信息3"},
				{Key: "/web/kkkkk3", Value: "一些信息4"},
				{Key: "/web/kkkkk4", Value: "一些信息5"},
			}

			go func() {
				for _, pair := range pairs {
					if leasePair, err := register.RegisterPair(pair, 40); err != nil {
						register.Logger.Err(err).Msg("[ETCD]")
					} else {
						// 如果需要挂载写入状态，可以读 keepAlive 通道。
						go func(pair *ext.EtcdPair) {
							for resp := range leasePair.KeepAliveChan {
								register.Logger.Info().
									Any("keep", resp).
									Any("pair", pair).
									Msg("[ETCD]")
							}
							register.Logger.Info().
								Any("keep", "关闭").
								Any("pair", pair).
								Msg("[ETCD]")
						}(&pair)
					}
				}
			}()
			return nil
		}); err != nil {
			return err
		}

		// ETCD 发现服务
		return container.Invoke(func(discovery *ext.EtcdDiscovery) error {
			discovery.Logger.Info().Str("action", "发现服务...").Msg("[ETCD]")
			go func() {
				if err := discovery.WatchService("/web"); err != nil {
					discovery.Logger.Err(err).Msg("[ETCD]")
				}
			}()
			go ext.Tick(10*time.Second, func() error {
				discovery.Logger.Info().Any("list", discovery.ListService()).Msg("[ETCD]")
				return nil
			})
			return nil
		})
	}
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
