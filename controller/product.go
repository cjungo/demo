package controller

import (
	"fmt"
	"time"

	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/cjungo/mid"
	localModel "github.com/cjungo/demo/local/model"
	"github.com/cjungo/demo/misc"
	"github.com/cjungo/demo/model"
	"github.com/rs/zerolog"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type ProductController struct {
	sqlite       *db.Sqlite
	mysql        *db.MySql
	logger       *zerolog.Logger
	tokenManager *mid.PermitManager[int32, misc.EmployeeToken]
}

// 本示例因为只是为了展示可选功能，所以让 MYSQL 可选，正常项目不会让数据库可选
type ProductControllerDi struct {
	dig.In
	Sqlite       *db.Sqlite
	TokenManager *mid.PermitManager[int32, misc.EmployeeToken]
	Logger       *zerolog.Logger
	MySql        *db.MySql `optional:"true"`
}

func NewProductController(di ProductControllerDi) (*ProductController, error) {
	return &ProductController{
		sqlite:       di.Sqlite,
		mysql:        di.MySql,
		logger:       di.Logger,
		tokenManager: di.TokenManager,
	}, nil
}

func (controller *ProductController) Add(ctx cjungo.HttpContext) (err error) {
	if controller.mysql != nil {
		m := &model.CjProduct{}
		if err = ctx.Bind(&m); err != nil {
			return ctx.RespBad(err)
		}
		tx := controller.mysql.Begin()
		if err = tx.Error; err != nil {
			return ctx.RespBad(err)
		}
		defer func() {
			if err != nil {
				tx.Rollback()
			}
		}()

		ltx := controller.sqlite.Begin()
		if err = ltx.Error; err != nil {
			return ctx.RespBad(err)
		}
		defer func() {
			if err != nil {
				ltx.Rollback()
			}
		}()

		if err = tx.Create(m).Error; err != nil {
			return ctx.RespBad(err)
		}
		mo := &localModel.CjOperation{
			OperatorID:     int32(m.ID),
			OperateAt:      time.Now(),
			OperateSummary: "添加样品",
		}
		if err = ltx.Create(mo).Error; err != nil {
			return ctx.RespBad(err)
		}
		if err = tx.Commit().Error; err != nil {
			return ctx.RespBad(err)
		}
		if err = ltx.Commit().Error; err != nil {
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
			return controller.sqlite.Transaction(func(ltx *gorm.DB) error {
				reqID := ctx.GetReqID()
				e := &misc.EmployeeToken{}
				if b := controller.tokenManager.GetToken(reqID, e); !b {
					return fmt.Errorf("无效的 TOKEN ID %s", reqID)
				}
				now := time.Now()
				m.CreateBy = e.EmployeeId
				m.CreateAt = now
				m.UpdateBy = e.EmployeeId
				m.UpdateAt = now
				if err := tx.Save(m).Error; err != nil {
					return err
				}

				mo := &localModel.CjOperation{
					OperatorID:     e.EmployeeId,
					OperateAt:      now,
					OperateSummary: "修改样品",
				}
				if err := ltx.Create(mo).Error; err != nil {
					return err
				}
				return nil
			})
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
