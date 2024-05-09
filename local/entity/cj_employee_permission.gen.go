// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/cjungo/demo/local/model"
)

func newCjEmployeePermission(db *gorm.DB, opts ...gen.DOOption) cjEmployeePermission {
	_cjEmployeePermission := cjEmployeePermission{}

	_cjEmployeePermission.cjEmployeePermissionDo.UseDB(db, opts...)
	_cjEmployeePermission.cjEmployeePermissionDo.UseModel(&model.CjEmployeePermission{})

	tableName := _cjEmployeePermission.cjEmployeePermissionDo.TableName()
	_cjEmployeePermission.ALL = field.NewAsterisk(tableName)
	_cjEmployeePermission.ID = field.NewInt32(tableName, "id")
	_cjEmployeePermission.EmployeeID = field.NewInt32(tableName, "employee_id")
	_cjEmployeePermission.PermissionID = field.NewInt32(tableName, "permission_id")

	_cjEmployeePermission.fillFieldMap()

	return _cjEmployeePermission
}

type cjEmployeePermission struct {
	cjEmployeePermissionDo cjEmployeePermissionDo

	ALL          field.Asterisk
	ID           field.Int32
	EmployeeID   field.Int32
	PermissionID field.Int32

	fieldMap map[string]field.Expr
}

func (c cjEmployeePermission) Table(newTableName string) *cjEmployeePermission {
	c.cjEmployeePermissionDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cjEmployeePermission) As(alias string) *cjEmployeePermission {
	c.cjEmployeePermissionDo.DO = *(c.cjEmployeePermissionDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cjEmployeePermission) updateTableName(table string) *cjEmployeePermission {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt32(table, "id")
	c.EmployeeID = field.NewInt32(table, "employee_id")
	c.PermissionID = field.NewInt32(table, "permission_id")

	c.fillFieldMap()

	return c
}

func (c *cjEmployeePermission) WithContext(ctx context.Context) *cjEmployeePermissionDo {
	return c.cjEmployeePermissionDo.WithContext(ctx)
}

func (c cjEmployeePermission) TableName() string { return c.cjEmployeePermissionDo.TableName() }

func (c cjEmployeePermission) Alias() string { return c.cjEmployeePermissionDo.Alias() }

func (c cjEmployeePermission) Columns(cols ...field.Expr) gen.Columns {
	return c.cjEmployeePermissionDo.Columns(cols...)
}

func (c *cjEmployeePermission) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cjEmployeePermission) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 3)
	c.fieldMap["id"] = c.ID
	c.fieldMap["employee_id"] = c.EmployeeID
	c.fieldMap["permission_id"] = c.PermissionID
}

func (c cjEmployeePermission) clone(db *gorm.DB) cjEmployeePermission {
	c.cjEmployeePermissionDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cjEmployeePermission) replaceDB(db *gorm.DB) cjEmployeePermission {
	c.cjEmployeePermissionDo.ReplaceDB(db)
	return c
}

type cjEmployeePermissionDo struct{ gen.DO }

func (c cjEmployeePermissionDo) Debug() *cjEmployeePermissionDo {
	return c.withDO(c.DO.Debug())
}

func (c cjEmployeePermissionDo) WithContext(ctx context.Context) *cjEmployeePermissionDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cjEmployeePermissionDo) ReadDB() *cjEmployeePermissionDo {
	return c.Clauses(dbresolver.Read)
}

func (c cjEmployeePermissionDo) WriteDB() *cjEmployeePermissionDo {
	return c.Clauses(dbresolver.Write)
}

func (c cjEmployeePermissionDo) Session(config *gorm.Session) *cjEmployeePermissionDo {
	return c.withDO(c.DO.Session(config))
}

func (c cjEmployeePermissionDo) Clauses(conds ...clause.Expression) *cjEmployeePermissionDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cjEmployeePermissionDo) Returning(value interface{}, columns ...string) *cjEmployeePermissionDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cjEmployeePermissionDo) Not(conds ...gen.Condition) *cjEmployeePermissionDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cjEmployeePermissionDo) Or(conds ...gen.Condition) *cjEmployeePermissionDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cjEmployeePermissionDo) Select(conds ...field.Expr) *cjEmployeePermissionDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cjEmployeePermissionDo) Where(conds ...gen.Condition) *cjEmployeePermissionDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cjEmployeePermissionDo) Order(conds ...field.Expr) *cjEmployeePermissionDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cjEmployeePermissionDo) Distinct(cols ...field.Expr) *cjEmployeePermissionDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cjEmployeePermissionDo) Omit(cols ...field.Expr) *cjEmployeePermissionDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cjEmployeePermissionDo) Join(table schema.Tabler, on ...field.Expr) *cjEmployeePermissionDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cjEmployeePermissionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *cjEmployeePermissionDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cjEmployeePermissionDo) RightJoin(table schema.Tabler, on ...field.Expr) *cjEmployeePermissionDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cjEmployeePermissionDo) Group(cols ...field.Expr) *cjEmployeePermissionDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cjEmployeePermissionDo) Having(conds ...gen.Condition) *cjEmployeePermissionDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cjEmployeePermissionDo) Limit(limit int) *cjEmployeePermissionDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cjEmployeePermissionDo) Offset(offset int) *cjEmployeePermissionDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cjEmployeePermissionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *cjEmployeePermissionDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cjEmployeePermissionDo) Unscoped() *cjEmployeePermissionDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cjEmployeePermissionDo) Create(values ...*model.CjEmployeePermission) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cjEmployeePermissionDo) CreateInBatches(values []*model.CjEmployeePermission, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cjEmployeePermissionDo) Save(values ...*model.CjEmployeePermission) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cjEmployeePermissionDo) First() (*model.CjEmployeePermission, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjEmployeePermission), nil
	}
}

func (c cjEmployeePermissionDo) Take() (*model.CjEmployeePermission, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjEmployeePermission), nil
	}
}

func (c cjEmployeePermissionDo) Last() (*model.CjEmployeePermission, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjEmployeePermission), nil
	}
}

func (c cjEmployeePermissionDo) Find() ([]*model.CjEmployeePermission, error) {
	result, err := c.DO.Find()
	return result.([]*model.CjEmployeePermission), err
}

func (c cjEmployeePermissionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CjEmployeePermission, err error) {
	buf := make([]*model.CjEmployeePermission, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cjEmployeePermissionDo) FindInBatches(result *[]*model.CjEmployeePermission, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cjEmployeePermissionDo) Attrs(attrs ...field.AssignExpr) *cjEmployeePermissionDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cjEmployeePermissionDo) Assign(attrs ...field.AssignExpr) *cjEmployeePermissionDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cjEmployeePermissionDo) Joins(fields ...field.RelationField) *cjEmployeePermissionDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cjEmployeePermissionDo) Preload(fields ...field.RelationField) *cjEmployeePermissionDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cjEmployeePermissionDo) FirstOrInit() (*model.CjEmployeePermission, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjEmployeePermission), nil
	}
}

func (c cjEmployeePermissionDo) FirstOrCreate() (*model.CjEmployeePermission, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjEmployeePermission), nil
	}
}

func (c cjEmployeePermissionDo) FindByPage(offset int, limit int) (result []*model.CjEmployeePermission, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cjEmployeePermissionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cjEmployeePermissionDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cjEmployeePermissionDo) Delete(models ...*model.CjEmployeePermission) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cjEmployeePermissionDo) withDO(do gen.Dao) *cjEmployeePermissionDo {
	c.DO = *do.(*gen.DO)
	return c
}
