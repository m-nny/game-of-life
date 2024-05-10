package main

import (
	"minmax.uk/game-of-life/pkg/engine/halflife"
)

func main() {
	level := 2
	u := halflife.BuildUniverse(level)

	// for row := range 4 {
	// 	for col := range 4 {
	// 		u.Set(row, col, true)
	// 	}
	// }
	u.Set(0, 1, true)
	u.Set(1, 1, true)
	u.Set(2, 1, true)
	u.DebugPrint()

	u.Iterate()
	u.DebugPrint()
}
