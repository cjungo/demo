package controller

import (
	"time"

	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/demo/model"
	"go.uber.org/dig"
	"gorm.io/gorm"
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
		tx := controller.mysql.Begin()
		if tx.Error != nil {
			return ctx.RespBad(tx.Error)
		}
		if err := tx.Create(m).Error; err != nil {
			tx.Rollback()
			return ctx.RespBad(err)
		}
		mo := &model.CjProductOperation{
			ProductID:   m.ID,
			OperateAt:   time.Now(),
			OperateType: 1,
		}
		if err := tx.Create(mo).Error; err != nil {
			tx.Rollback()
			return ctx.RespBad(err)
		}
		if err := tx.Commit().Error; err != nil {
			return ctx.RespBad(err)
		}
		return ctx.Resp(m)
	} else {
		return ctx.Resp("没有数据库")
	}
}

func (controller *ProductController) Edit(ctx cjungo.HttpContext) error {
	if controller.mysql != nil {
		m := &model.CjProduct{}
		if err := ctx.Bind(&m); err != nil {
			return ctx.RespBad(err)
		}
		if err := controller.mysql.Transaction(func(tx *gorm.DB) error {
			if err := tx.Save(m).Error; err != nil {
				return err
			}
			mo := &model.CjProductOperation{
				ProductID:   m.ID,
				OperateAt:   time.Now(),
				OperateType: 3,
			}
			if err := tx.Create(mo).Error; err != nil {
				return err
			}
			return nil
		}); err != nil {
			return ctx.RespBad(err)
		}

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
