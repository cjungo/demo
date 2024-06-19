package main

// etcd 示例

import (
	"time"

	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/ext"
)

// 注册 etcd 服务，根据项目业务，把自身提供的功能注册到 etcd 。
func registerEtcdService(container cjungo.DiContainer) error {
	return container.Invoke(func(register *ext.EtcdRegister) error {
		pairs := []ext.EtcdPair{
			{Key: "/web/kkkkk1", Value: "一些信息2"},
			{Key: "/web/kkkkk2", Value: "一些信息3"},
			{Key: "/web/kkkkk3", Value: "一些信息4"},
			{Key: "/web/kkkkk4", Value: "一些信息5"},
		}

		// 此处若要确保 etcd 必须，可以去掉 go，因为示例，所以 etcd 非必须。
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
	})
}

// 根据 项目 的需求，把依赖的 etcd 服务发现出来。
func registerEtcdDiscovery(container cjungo.DiContainer) error {
	return container.Invoke(func(discovery *ext.EtcdDiscovery) error {
		discovery.Logger.Info().Str("action", "发现服务...").Msg("[ETCD]")

		// 此处若要确保 etcd 必须，可以去掉 go，因为示例，所以 etcd 非必须。
		// 建议 必须，一般项目依赖其他项目都是强依赖，需要其他服务可用才能正常运行。
		go func() {
			if err := discovery.WatchService("/web"); err != nil {
				discovery.Logger.Err(err).Msg("[ETCD]")
			}
		}()

		// 日志打印 etcd 状态，可选
		go ext.Tick(10*time.Second, func() error {
			discovery.Logger.Info().Any("list", discovery.ListService()).Msg("[ETCD]")
			return nil
		})
		return nil
	})
}
