package main

import (
	"log"

	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/demo/cmds"
	"gorm.io/gorm"
)

func main() {
	if err := cjungo.RunCommand[any](
		func(mysql *db.MySql) error {
			return mysql.TransactionSilent(func(d *gorm.DB) error {
				return d.Exec("SHOW DATABASES").Error
			})
		},
		db.LoadMySqlConfFormEnv,
		db.LoadSqliteConfFormEnv,
		cmds.InitMySql(),
		cmds.InitSqlite(),
	); err != nil {
		log.Fatalln(err)
	}
}
