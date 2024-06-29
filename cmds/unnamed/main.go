package main

import (
	"log"

	"github.com/cjungo/cjungo"
	"github.com/rs/zerolog"
)

func main() {
	// 使用闭包时，日志文件和任务名会变成目录名。
	if err := cjungo.RunCommand(func(logger *zerolog.Logger) {
		logger.
			Info().
			Str("command", "unnamed").
			Msg("[CMD]")
	}); err != nil {
		log.Fatalln(err)
	}
}
