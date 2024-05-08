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

	"github.com/cjungo/demo/model"
)

func newCjProduct(db *gorm.DB, opts ...gen.DOOption) cjProduct {
	_cjProduct := cjProduct{}

	_cjProduct.cjProductDo.UseDB(db, opts...)
	_cjProduct.cjProductDo.UseModel(&model.CjProduct{})

	tableName := _cjProduct.cjProductDo.TableName()
	_cjProduct.ALL = field.NewAsterisk(tableName)
	_cjProduct.ID = field.NewUint32(tableName, "id")
	_cjProduct.Number = field.NewString(tableName, "number")
	_cjProduct.Shortname = field.NewString(tableName, "shortname")
	_cjProduct.Fullname = field.NewString(tableName, "fullname")
	_cjProduct.IsRemoved = field.NewUint32(tableName, "is_removed")

	_cjProduct.fillFieldMap()

	return _cjProduct
}

// cjProduct 产品
type cjProduct struct {
	cjProductDo cjProductDo

	ALL       field.Asterisk
	ID        field.Uint32
	Number    field.String // 编号
	Shortname field.String
	Fullname  field.String
	IsRemoved field.Uint32

	fieldMap map[string]field.Expr
}

func (c cjProduct) Table(newTableName string) *cjProduct {
	c.cjProductDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cjProduct) As(alias string) *cjProduct {
	c.cjProductDo.DO = *(c.cjProductDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cjProduct) updateTableName(table string) *cjProduct {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint32(table, "id")
	c.Number = field.NewString(table, "number")
	c.Shortname = field.NewString(table, "shortname")
	c.Fullname = field.NewString(table, "fullname")
	c.IsRemoved = field.NewUint32(table, "is_removed")

	c.fillFieldMap()

	return c
}

func (c *cjProduct) WithContext(ctx context.Context) *cjProductDo {
	return c.cjProductDo.WithContext(ctx)
}

func (c cjProduct) TableName() string { return c.cjProductDo.TableName() }

func (c cjProduct) Alias() string { return c.cjProductDo.Alias() }

func (c cjProduct) Columns(cols ...field.Expr) gen.Columns { return c.cjProductDo.Columns(cols...) }

func (c *cjProduct) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cjProduct) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 5)
	c.fieldMap["id"] = c.ID
	c.fieldMap["number"] = c.Number
	c.fieldMap["shortname"] = c.Shortname
	c.fieldMap["fullname"] = c.Fullname
	c.fieldMap["is_removed"] = c.IsRemoved
}

func (c cjProduct) clone(db *gorm.DB) cjProduct {
	c.cjProductDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cjProduct) replaceDB(db *gorm.DB) cjProduct {
	c.cjProductDo.ReplaceDB(db)
	return c
}

type cjProductDo struct{ gen.DO }

func (c cjProductDo) Debug() *cjProductDo {
	return c.withDO(c.DO.Debug())
}

func (c cjProductDo) WithContext(ctx context.Context) *cjProductDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cjProductDo) ReadDB() *cjProductDo {
	return c.Clauses(dbresolver.Read)
}

func (c cjProductDo) WriteDB() *cjProductDo {
	return c.Clauses(dbresolver.Write)
}

func (c cjProductDo) Session(config *gorm.Session) *cjProductDo {
	return c.withDO(c.DO.Session(config))
}

func (c cjProductDo) Clauses(conds ...clause.Expression) *cjProductDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cjProductDo) Returning(value interface{}, columns ...string) *cjProductDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cjProductDo) Not(conds ...gen.Condition) *cjProductDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cjProductDo) Or(conds ...gen.Condition) *cjProductDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cjProductDo) Select(conds ...field.Expr) *cjProductDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cjProductDo) Where(conds ...gen.Condition) *cjProductDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cjProductDo) Order(conds ...field.Expr) *cjProductDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cjProductDo) Distinct(cols ...field.Expr) *cjProductDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cjProductDo) Omit(cols ...field.Expr) *cjProductDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cjProductDo) Join(table schema.Tabler, on ...field.Expr) *cjProductDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cjProductDo) LeftJoin(table schema.Tabler, on ...field.Expr) *cjProductDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cjProductDo) RightJoin(table schema.Tabler, on ...field.Expr) *cjProductDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cjProductDo) Group(cols ...field.Expr) *cjProductDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cjProductDo) Having(conds ...gen.Condition) *cjProductDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cjProductDo) Limit(limit int) *cjProductDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cjProductDo) Offset(offset int) *cjProductDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cjProductDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *cjProductDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cjProductDo) Unscoped() *cjProductDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cjProductDo) Create(values ...*model.CjProduct) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cjProductDo) CreateInBatches(values []*model.CjProduct, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cjProductDo) Save(values ...*model.CjProduct) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cjProductDo) First() (*model.CjProduct, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjProduct), nil
	}
}

func (c cjProductDo) Take() (*model.CjProduct, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjProduct), nil
	}
}

func (c cjProductDo) Last() (*model.CjProduct, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjProduct), nil
	}
}

func (c cjProductDo) Find() ([]*model.CjProduct, error) {
	result, err := c.DO.Find()
	return result.([]*model.CjProduct), err
}

func (c cjProductDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CjProduct, err error) {
	buf := make([]*model.CjProduct, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cjProductDo) FindInBatches(result *[]*model.CjProduct, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cjProductDo) Attrs(attrs ...field.AssignExpr) *cjProductDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cjProductDo) Assign(attrs ...field.AssignExpr) *cjProductDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cjProductDo) Joins(fields ...field.RelationField) *cjProductDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cjProductDo) Preload(fields ...field.RelationField) *cjProductDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cjProductDo) FirstOrInit() (*model.CjProduct, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjProduct), nil
	}
}

func (c cjProductDo) FirstOrCreate() (*model.CjProduct, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjProduct), nil
	}
}

func (c cjProductDo) FindByPage(offset int, limit int) (result []*model.CjProduct, count int64, err error) {
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

func (c cjProductDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cjProductDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cjProductDo) Delete(models ...*model.CjProduct) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cjProductDo) withDO(do gen.Dao) *cjProductDo {
	c.DO = *do.(*gen.DO)
	return c
}
