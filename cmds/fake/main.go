package main

import (
	"log"

	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/demo/entity"
	localEntity "github.com/cjungo/demo/local/entity"
	localModel "github.com/cjungo/demo/local/model"
	"github.com/cjungo/demo/misc"
	"gorm.io/gorm"
)

type FakeEmployeeCommand struct {
}

func (command *FakeEmployeeCommand) Init(c cjungo.DiContainer) error {
	// 注册数据库
	if err := c.Provide(db.NewMySqlHandle(func(mysql *db.MySql) error {
		entity.Use(mysql.DB)
		return nil
	})); err != nil {
		return err
	}

	if err := c.Provide(db.NewSqliteHandle(func(sqlite *db.Sqlite) error {
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
	})); err != nil {
		return err
	}
	return nil
}

func (command *FakeEmployeeCommand) Exec(c cjungo.DiContainer) error {
	return nil
}

func init() {
	if err := cjungo.LoadEnv(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	command, err := cjungo.NewCommand(&FakeEmployeeCommand{})
	if err != nil {
		log.Fatalln(err)
	}

	err = command.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
