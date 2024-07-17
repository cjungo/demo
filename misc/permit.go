package misc

import (
	"time"

	"github.com/cjungo/cjungo/ext"
	localModel "github.com/cjungo/demo/local/model"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elliotchance/pie/v2"
	"gorm.io/gorm"
)

var Permissions []localModel.CjPermission

func init() {
	Permissions = []localModel.CjPermission{
		{ID: 10000, ParentID: 0, Tag: "default", Name: "基础权限"},
		{ID: 11000, ParentID: 0, Tag: "employee", Name: "员工管理"},
		{ID: 11001, ParentID: 11000, Tag: "employee_find", Name: "员工查看"},
		{ID: 11002, ParentID: 11000, Tag: "employee_add", Name: "员工添加"},
		{ID: 11003, ParentID: 11000, Tag: "employee_edit", Name: "员工修改"},
		{ID: 12000, ParentID: 0, Tag: "product", Name: "样品管理"},
		{ID: 12001, ParentID: 12000, Tag: "product_find", Name: "样品查看"},
		{ID: 12002, ParentID: 12000, Tag: "product_add", Name: "样品添加"},
		{ID: 12003, ParentID: 12000, Tag: "product_edit", Name: "样品修改"},
	}
}

func EnsurePermissions(tx *gorm.DB) error {
	permissions := []localModel.CjPermission{}
	if err := tx.Find(&permissions).Error; err != nil {
		return err
	}
	added, removed := pie.Diff(permissions, Permissions)
	if len(removed) > 0 {
		if err := tx.Delete(removed).Error; err != nil {
			return err
		}
	}

	if len(added) > 0 {
		if err := tx.CreateInBatches(added, 100).Error; err != nil {
			return err
		}
	}
	return nil
}

func EnsureAdmin(tx *gorm.DB) error {
	admin := &localModel.CjEmployee{}
	if err := tx.Find(admin, 1).Error; err != nil {
		return err
	}
	if admin.ID != 1 {
		now := time.Now()
		admin = &localModel.CjEmployee{
			ID:       1,
			Username: "admin",
			Password: ext.Sha256("admin").Hex(),
			Nickname: "admin",
			CreateBy: 0,
			CreateAt: now,
			UpdateBy: 0,
			UpdateAt: now,
		}
		if err := tx.Save(admin).Error; err != nil {
			return err
		}
	}
	return EnsureEmployeePermissions(tx, admin)
}

func EnsureEmployeePermissions(tx *gorm.DB, employee *localModel.CjEmployee) error {
	epids := []int32{}
	if err := tx.Table("cj_employee_permission").Select("permission_id").Where("employee_id=?", employee.ID).Find(&epids).Error; err != nil {
		return err
	}
	epidset := mapset.NewSet(epids...)
	needs := pie.Filter(Permissions, func(p localModel.CjPermission) bool {
		return !epidset.ContainsOne(p.ID)
	})
	if len(needs) > 0 {
		added := make([]localModel.CjEmployeePermission, len(needs))
		for i, p := range needs {
			added[i].EmployeeID = employee.ID
			added[i].PermissionID = p.ID
		}
		if err := tx.Save(added).Error; err != nil {
			return err
		}
	}
	return nil
}
