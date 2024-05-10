package engine

import (
	"fmt"

	"minmax.uk/game-of-life/pkg/boards"
)

type Engine interface {
	fmt.Stringer
	Iterate()
	Name() string
}

type EngineBuilder func(boards.BoardSpec) (Engine, error)
