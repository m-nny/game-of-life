package main

import (
	"fmt"

	"minmax.uk/game-of-life/pkg/engine/halflife"
)

func main() {
	// for level := range 3 {
	// 	u := halflife.BuildUniverse(level)
	// 	u.DebugPrint()
	// }
	level := 2
	u := halflife.BuildUniverse(level)
	u.DebugPrint()
	fmt.Println()

	for i := range 1 << level {
		fmt.Printf("Set %d %d\n", i, i)
		u.Set(i, i, true)
		u.DebugPrint()
	}

}
