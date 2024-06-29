package main

import (
	"log"

	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/demo/cmds"
	"github.com/rs/zerolog"
)

func FakeEmployee(logger *zerolog.Logger, mysql *db.MySql) error {
	logger.
		Info().
		Str("command", "FakeEmployee").
		Msg("[CMD]")
	return nil
}

func main() {
	// 使用命名函数时，日志文件和任务名是函数名。
	if err := cjungo.RunCommand(
		FakeEmployee,
		db.LoadMySqlConfFormEnv,
		db.LoadSqliteConfFormEnv,
		cmds.InitMySql(),
		cmds.InitSqlite(),
	); err != nil {
		log.Fatalln(err)
	}
}
