package controller

import (
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/demo/local/model"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type EmployeeController struct {
	logger *zerolog.Logger
	sqlite *db.Sqlite
}

func NewEmployeeController(
	logger *zerolog.Logger,
	sqlite *db.Sqlite,
) *EmployeeController {
	return &EmployeeController{
		logger: logger,
		sqlite: sqlite,
	}
}

// 此种方式，会使得参数多出 ID 这种不应该出现的字段。
// 虽然可以不传，却脏。
type EmployeeAddParam struct {
	model.CjEmployee
	Permissions []int32 `json:"permissions" form:"permissions"`
}

func (controller *EmployeeController) Add(ctx cjungo.HttpContext) error {
	param := &EmployeeAddParam{}
	if err := ctx.Bind(param); err != nil {
		return ctx.RespBad(err)
	}

	if err := controller.sqlite.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&param.CjEmployee).Error; err != nil {
			return err
		}

		permissions := make([]model.CjEmployeePermission, len(param.Permissions))
		for i, pid := range param.Permissions {
			permissions[i] = model.CjEmployeePermission{
				EmployeeID:   param.ID,
				PermissionID: pid,
			}
		}

		if err := tx.CreateInBatches(permissions, 100).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return ctx.Resp(err)
	}

	return ctx.Resp(param)
}

type EmployeeEditParam EmployeeAddParam

func (controller *EmployeeController) Edit(ctx cjungo.HttpContext) error {
	param := &EmployeeEditParam{}
	if err := ctx.Bind(param); err != nil {
		return ctx.RespBad(err)
	}

	// 推荐的事务使用， Transaction 要比 Begin 方便。
	if err := controller.sqlite.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&param.CjEmployee).Error; err != nil {
			return err
		}
		if err := tx.Where("employee_id=?", param.ID).
			Delete(&model.CjEmployeePermission{}).Error; err != nil {
			return err
		}

		permissions := make([]model.CjEmployeePermission, len(param.Permissions))
		for i, pid := range param.Permissions {
			permissions[i] = model.CjEmployeePermission{
				EmployeeID:   param.ID,
				PermissionID: pid,
			}
		}
		if err := tx.CreateInBatches(permissions, 100).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return ctx.Resp(err)
	}

	return ctx.Resp(param)
}

type EmployeeDetailResult EmployeeAddParam

func (controller *EmployeeController) Detail(ctx cjungo.HttpContext) error {
	id := ctx.QueryParam("id")
	result := &EmployeeDetailResult{}
	if err := controller.sqlite.First(&result.CjEmployee, id).Error; err != nil {
		return ctx.RespBad(err)
	}

	if err := controller.sqlite.Select("permission_id").
		Table("cj_employee_permission").
		Where("employee_id=?", result.ID).
		Find(&result.Permissions).Error; err != nil {
		return ctx.RespBad(err)
	}

	return ctx.Resp(result)
}

func (controller *EmployeeController) Delete(ctx cjungo.HttpContext) error {
	id := ctx.QueryParam("id")
	if err := controller.sqlite.Delete(&model.CjEmployee{}, id).Error; err != nil {
		return ctx.RespBad(err)
	}
	return ctx.RespOk()
}
