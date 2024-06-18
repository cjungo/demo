package controller

import (
	"fmt"
	"time"

	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/cjungo/ext"
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
	tokenManager *mid.PermitManager[string, misc.EmployeeToken]
}

// 本示例因为只是为了展示可选功能，所以让 MYSQL 可选，正常项目不会让数据库可选
type ProductControllerDi struct {
	dig.In
	Sqlite       *db.Sqlite
	TokenManager *mid.PermitManager[string, misc.EmployeeToken]
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

// 使用 MoveField 可以缩小暴露接口的参数，但是麻烦。
type ProductAddParam struct {
	Number    string  `json:"number"` // 编号
	Shortname *string `json:"shortname"`
	Fullname  *string `json:"fullname"`
}

func (controller *ProductController) Add(ctx cjungo.HttpContext) error {
	if controller.mysql != nil {
		param := &ProductAddParam{}
		if err := ctx.Bind(&param); err != nil {
			return ctx.RespBad(err)
		}

		pp, ok := controller.tokenManager.GetProof(ctx)
		if !ok {
			return fmt.Errorf("无效TOKEN ID")
		}
		token := pp.GetToken()

		now := time.Now()
		m := &model.CjProduct{
			CreateBy: token.EmployeeId,
			CreateAt: now,
			UpdateBy: token.EmployeeId,
			UpdateAt: now,
		}
		ext.MoveField(param, m)

		controller.logger.Info().Any("product", m).Msg("ProductAdd")

		// 这里只是示例可以使用事务的一种方式。Begin 这种形式比较底层，最后阶段比较麻烦。
		// 建议还是使用 Transaction 的形式。参考 employee 控制器。
		tx := controller.mysql.Begin()
		if err := tx.Error; err != nil {
			return ctx.RespBad(err)
		}
		txc := make(chan bool, 1)
		defer func() {
			select {
			case <-txc:
				controller.logger.Info().Msg("ProductAdd")
			default:
				controller.logger.Error().Msg("ProductAdd")
				tx.Rollback()
			}
		}()

		ltx := controller.sqlite.Begin()
		if err := ltx.Error; err != nil {
			return ctx.RespBad(err)
		}
		ltxc := make(chan bool, 1)
		defer func() {
			select {
			case <-ltxc:
				controller.logger.Info().Msg("ProductAdd")
			default:
				controller.logger.Error().Msg("ProductAdd")
				ltx.Rollback()
			}
		}()

		controller.logger.Info().Str("tip", "事务START").Msg("ProductAdd")

		if err := tx.Create(m).Error; err != nil {
			return ctx.RespBad(err)
		}
		mo := &localModel.CjOperation{
			OperatorID:     int32(m.ID),
			OperateAt:      time.Now(),
			OperateSummary: "添加样品",
		}
		if err := ltx.Create(mo).Error; err != nil {
			return ctx.RespBad(err)
		}
		controller.logger.Info().Str("tip", "事务END").Msg("ProductAdd")
		if err := tx.Commit().Error; err != nil {
			return ctx.RespBad(err)
		}
		if err := ltx.Commit().Error; err != nil {
			return ctx.RespBad(err)
		}
		txc <- true
		ltxc <- true
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
				pp, b := controller.tokenManager.GetProof(ctx)
				if !b {
					return fmt.Errorf("无效的 TOKEN ID %s", ctx.GetReqID())
				}
				e := pp.GetToken()
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
