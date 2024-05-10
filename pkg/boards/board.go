package boards

import (
	"math/rand"
	"strings"

	"minmax.uk/game-of-life/pkg/utils"
)

type BoardSpec struct {
	Name string
	Rows int64
	Cols int64
	Str  string
	Freq int
}

func (b *BoardSpec) Unpack64() (rows, cols int64, str string) {
	return b.Rows, b.Cols, b.Str
}

func (b *BoardSpec) Normalized() string {
	return NormalizeBoard(b.Str)
}

var Block = BoardSpec{
	Name: "Block",
	Rows: 4,
	Cols: 4,
	Freq: 0,
	Str: `
		....
		.##.
		.##.
		....
		`,
}

var Blinker = BoardSpec{
	Name: "Bliner",
	Rows: 5,
	Cols: 5,
	Freq: 2,
	Str: `
		.....
		..#..
		..#..
		..#..
		.....
		`,
}

var Beehive = BoardSpec{
	Name: "Beehive",
	Rows: 5,
	Cols: 6,
	Freq: 0,
	Str: `
		......
		..##..
		.#..#.
		..##..
		......
		`,
}

var Toad = BoardSpec{
	Name: "Toad",
	Rows: 6,
	Cols: 6,
	Freq: 2,
	Str: `
		......
		...#..
		.#..#.
		.#..#.
		..#...
		......
		`,
}

var Pulsar = BoardSpec{
	Name: "Pulsar",
	Rows: 17,
	Cols: 17,
	Freq: 3,
	Str: `
		.................
		.................
		....OOO...OOO....
		.................
		..O....O.O....O..
		..O....O.O....O..
		..O....O.O....O..
		....OOO...OOO....
		.................
		....OOO...OOO....
		..O....O.O....O..
		..O....O.O....O..
		..O....O.O....O..
		.................
		....OOO...OOO....
		.................
		.................
		`,
}

func Random(rows, cols, seed int64) BoardSpec {
	utils.Assert(rows > 0 && cols > 0, "both rows and cols should be > 0")

	// TODO: add options for randomness
	rand := rand.New(rand.NewSource(seed))

	// We can do it more efficiently by generating whole chunks (int64) and using it
	var builder strings.Builder
	for range rows * cols {
		builder.WriteString(CELL_TO_CHAR[getRandBool(rand)])
	}
	return BoardSpec{
		Rows: rows,
		Cols: cols,
		Str:  builder.String(),
	}
}

var RAND_MAX_N = 100

func getRandBool(rand *rand.Rand) bool {
	return (rand.Intn(RAND_MAX_N) & 1) == 1
}
