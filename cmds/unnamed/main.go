package main

import (
	"log"

	"github.com/cjungo/cjungo"
	"github.com/rs/zerolog"
)

type Opts struct {
	// 示例名
	Name   string `short:"n" long:"name" description:"A name" required:"true"`
	Animal string `long:"animal" choice:"cat" choice:"dog"`
}

func main() {
	// 使用闭包时，日志文件和任务名会变成目录名。
	if err := cjungo.RunCommand[Opts](func(logger *zerolog.Logger, opts *Opts) {
		logger.
			Info().
			Str("command", "unnamed").
			Msg("[CMD]")
	}); err != nil {
		log.Fatalln(err)
	}
}
