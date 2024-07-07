package main

import (
	"log"

	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/ext"
	"github.com/rs/zerolog"
)

type SSA struct {
	Aint int16
}

type SA struct {
	Afloat float32
	Ssa    SSA
}

type SB struct {
	Afloat float64
	Ssa    SSA
}

func main() {
	if err := cjungo.RunCommand[any](func(logger *zerolog.Logger) {
		a := SA{
			Afloat: 1.234,
			Ssa:    SSA{Aint: 654},
		}
		b := SB{
			Afloat: 4.567,
			Ssa:    SSA{Aint: 987},
		}

		ext.MoveField(&a, &b)

		logger.Info().
			Any("a", a).
			Any("b", b).
			Msg("[MOVE FIELD]")

		ext.MoveFieldEx(&a, &b)

		logger.Info().
			Any("a", a).
			Any("b", b).
			Msg("[MOVE FIELD]")
	}); err != nil {
		log.Fatalln(err)
	}
}
