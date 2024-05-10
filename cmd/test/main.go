package main

import (
	"fmt"

	"minmax.uk/game-of-life/pkg/engine/halflife"
)

func main() {
	for level := range 3 {
		u := halflife.BuildUniverse(level)
		s := u.BoardString()
		fmt.Printf("level: %d\n", level)
		fmt.Printf("universe: %+v\n", u)
		fmt.Printf("%s\n", s)
		fmt.Println()
	}
}
