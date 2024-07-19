package misc

import (
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/ext"
)

type MyMessageToken string // 注意，没有等于号。
type MyMessageController = ext.MessageController[MyMessageToken]

func ProvideMyMessageController() ext.MessageControllerProvide[MyMessageToken] {
	return ext.ProvideMessageController(&ext.MessageControllerProviderConf[MyMessageToken]{
		TokenAccess: func(ctx cjungo.HttpContext) (MyMessageToken, error) {
			return MyMessageToken(ctx.GetReqID()), nil
		},
	})
}
