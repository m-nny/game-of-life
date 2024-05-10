package bitset_engine

import (
	"fmt"
	"strings"

	"minmax.uk/game-of-life/pkg/bitset"
	"minmax.uk/game-of-life/pkg/engine"
)

var _ engine.Engine = (*BitsetEngine)(nil)

type BitsetEngine struct {
	Rows int64
	Cols int64

	cells     []bool
	prevCells []bool
}

func (e *BitsetEngine) Name() string {
	return "BitsetEngine"
}

func (e *BitsetEngine) Iterate() {
	e.prevCells, e.cells = e.cells, e.prevCells
	for row := int64(1); row+1 < e.Rows; row++ {
		b := bitset.Empty()

		up_i := (row - 1) * e.Cols
		cur_i := row * e.Cols
		bot_i := (row + 1) * e.Cols

		b.SetForward(e.prevCells[up_i], e.prevCells[cur_i], e.prevCells[bot_i])
		up_i++
		cur_i++
		bot_i++
		b.Shift()

		b.SetForward(e.prevCells[up_i], e.prevCells[cur_i], e.prevCells[bot_i])
		up_i++
		cur_i++
		bot_i++
		b.Shift()

		for col := int64(1); col+1 < e.Cols; col++ {
			b.SetForward(e.prevCells[up_i], e.prevCells[cur_i], e.prevCells[bot_i])
			e.cells[cur_i-1] = bool(b.NextValue())

			up_i++
			cur_i++
			bot_i++
			b.Shift()
		}
	}
}

func debugString(item any) string {
	switch item := item.(type) {
	case bool:
		if item {
			return "#"
		} else {
			return "."
		}
	}
	return string(fmt.Sprintf("%v", item)[0])
}

func debugSlice[S ~[]E, E comparable](arr S, lineLen int64) string {
	var builder strings.Builder
	for i, item := range arr {
		builder.WriteString(debugString(item))
		if i%int(lineLen) == int(lineLen)-1 && i+1 != len(arr) {
			builder.WriteRune('\n')
		}
	}
	return builder.String()
}
