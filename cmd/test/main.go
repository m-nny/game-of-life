package main

import (
	"fmt"
	"log"
	"strings"

	"minmax.uk/game-of-life/pkg/bitset_engine"
	"minmax.uk/game-of-life/pkg/boards"
	"minmax.uk/game-of-life/pkg/engine/halflife"
)

var board = boards.BoardSpec{
	Rows: 16,
	Cols: 16,
	Str: `
	  ................
	  ................
	  ................
	  ................
		....OO....O.....
		....OO...O......
		.........OOO....
		................
		................
		....OOO..OO.....
		.........O.O....
		..........OO....
	  ................
	  ................
	  ................
	  ................
		`,
}

func run_halflife() (string, error) {
	u, err := halflife.FromBoardSpec(board)
	if err != nil {
		return "", err
	}

	u.DebugPrint(false)

	u.Iterate()
	u.DebugPrint(false)

	got := u.BoardString()

	return got, nil
}

func expect() (string, error) {
	g, err := bitset_engine.FromBoardSpec(board)
	if err != nil {
		return "", err
	}
	g.Iterate()
	lines := strings.Split(g.String(), "\n")
	start := board.Rows >> 2
	new_w := board.Rows >> 1
	lines = lines[start : start+new_w]
	for i := range lines {
		lines[i] = lines[i][start : start+new_w]
	}
	return strings.Join(lines, "\n"), nil
}

func main() {
	got, err := run_halflife()
	if err != nil {
		log.Fatal(err)
	}

	want, err := expect()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("want:\n%s\n", want)
	fmt.Printf("got:\n%s\n", got)
	if want != got {
		log.Fatal("want != got")
	}
	fmt.Printf("OK")
}
