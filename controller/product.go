package controller

import (
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/demo/model"
	"go.uber.org/dig"
)

type ProductController struct {
	mysql *db.MySql
}

// 本示例因为只是为了展示可选功能，所以让 MYSQL 可选，正常项目不会让数据库可选
type ProductControllerDi struct {
	dig.In
	MySql *db.MySql `optional:"true"`
}

func NewProductController(di ProductControllerDi) (*ProductController, error) {
	return &ProductController{
		mysql: di.MySql,
	}, nil
}

func (controller *ProductController) Add(ctx cjungo.HttpContext) error {
	if controller.mysql != nil {
		m := &model.CjProduct{}
		if err := ctx.Bind(&m); err != nil {
			return ctx.RespBad(err)
		}
		r := controller.mysql.Create(m)
		if r.Error != nil {
			return ctx.RespBad(r.Error)
		}
		// mo := &model.CjProductOperation{

		// }
		return ctx.Resp(m)
	} else {
		return ctx.Resp("没有数据库")
	}
}

func (controller *ProductController) Detail(ctx cjungo.HttpContext) error {
	if controller.mysql != nil {
		id := ctx.QueryParam("id")
		result := &model.CjProduct{}
		controller.mysql.Select("*").Where("id = ?", id).Find(result)
		return ctx.Resp(result)
	} else {
		return ctx.Resp("没有数据库")
	}
}
