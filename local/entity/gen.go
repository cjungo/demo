// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                   db,
		CjEmployee:           newCjEmployee(db, opts...),
		CjEmployeePermission: newCjEmployeePermission(db, opts...),
		CjOperation:          newCjOperation(db, opts...),
		CjPermission:         newCjPermission(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	CjEmployee           cjEmployee
	CjEmployeePermission cjEmployeePermission
	CjOperation          cjOperation
	CjPermission         cjPermission
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                   db,
		CjEmployee:           q.CjEmployee.clone(db),
		CjEmployeePermission: q.CjEmployeePermission.clone(db),
		CjOperation:          q.CjOperation.clone(db),
		CjPermission:         q.CjPermission.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                   db,
		CjEmployee:           q.CjEmployee.replaceDB(db),
		CjEmployeePermission: q.CjEmployeePermission.replaceDB(db),
		CjOperation:          q.CjOperation.replaceDB(db),
		CjPermission:         q.CjPermission.replaceDB(db),
	}
}

type queryCtx struct {
	CjEmployee           *cjEmployeeDo
	CjEmployeePermission *cjEmployeePermissionDo
	CjOperation          *cjOperationDo
	CjPermission         *cjPermissionDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		CjEmployee:           q.CjEmployee.WithContext(ctx),
		CjEmployeePermission: q.CjEmployeePermission.WithContext(ctx),
		CjOperation:          q.CjOperation.WithContext(ctx),
		CjPermission:         q.CjPermission.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
