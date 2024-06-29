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
		Msg("CMD")
	return nil
}

func main() {
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
