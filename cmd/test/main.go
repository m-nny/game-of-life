package main

import (
	"minmax.uk/game-of-life/pkg/engine/halflife"
)

func main() {
	for level := range 3 {
		u := halflife.BuildUniverse(level)
		u.DebugPrint()
	}
}
