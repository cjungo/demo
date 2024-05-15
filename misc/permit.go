package misc

import (
	localModel "github.com/cjungo/demo/local/model"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elliotchance/pie/v2"
	"gorm.io/gorm"
)

var Permissions []localModel.CjPermission

func init() {
	Permissions = []localModel.CjPermission{
		{ID: 12000, ParentID: 0, Tag: "product", Name: "样品"},
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
	added, _ := pie.Diff(permissions, Permissions)
	if len(added) > 0 {
		if err := tx.CreateInBatches(added, 100).Error; err != nil {
			return err
		}
	}
	return nil
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
	added := make([]localModel.CjEmployeePermission, len(needs))
	for i, p := range needs {
		added[i].EmployeeID = employee.ID
		added[i].PermissionID = p.ID
	}
	if err := tx.Save(added).Error; err != nil {
		return err
	}
	return nil
}
