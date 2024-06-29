package cmds

import (
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/demo/entity"
	localEntity "github.com/cjungo/demo/local/entity"
	localModel "github.com/cjungo/demo/local/model"
	"github.com/cjungo/demo/misc"
	"gorm.io/gorm"
)

// 获取 Mysql 初始化提供者
func InitMySql() db.MySqlProvide {
	return db.NewMySqlHandle(func(mysql *db.MySql) error {
		entity.Use(mysql.DB)
		return nil
	})
}

// 获取 Sqlite 初始化提供者
func InitSqlite() db.SqliteProvide {
	return db.NewSqliteHandle(func(sqlite *db.Sqlite) error {
		localEntity.Use(sqlite.DB)
		sqlite.AutoMigrate(
			&localModel.CjPermission{},
			&localModel.CjOperation{},
			&localModel.CjEmployee{},
			&localModel.CjEmployeePermission{},
		)
		return sqlite.Transaction(func(tx *gorm.DB) error {
			if err := misc.EnsurePermissions(tx); err != nil {
				return err
			}

			if err := misc.EnsureAdmin(tx); err != nil {
				return err
			}

			return nil
		})
	})
}
