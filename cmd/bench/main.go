package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"minmax.uk/game-of-life/pkg/bitset_engine"
	"minmax.uk/game-of-life/pkg/boards"
	"minmax.uk/game-of-life/pkg/engine"
	"minmax.uk/game-of-life/pkg/naive_engine"
)

var (
	rows  = flag.Int64("rows", 1000, "rows")
	cols  = flag.Int64("cols", 1000, "cols")
	seed  = flag.Int64("seed", 42, "seed")
	iters = flag.Int("iters", 1000, "iters")
	cnt   = flag.Int("cnt", 5, "number of times to run")

	engine_name = flag.String("engine", "bitset", "engine to benchmark")
)

func buildEngine(board boards.BoardSpec) (engine.Engine, error) {
	if *engine_name == "bitset" {
		return bitset_engine.FromBoardSpec(board)
	} else if *engine_name == "naive" {
		return naive_engine.FromBoardSpec(board)
	}
	return nil, fmt.Errorf("unknown engine %s", *engine_name)
}

func run(board boards.BoardSpec) (time.Duration, error) {
	g, err := buildEngine(board)
	if err != nil {
		return 0, err
	}
	start_time := time.Now()
	for range *iters {
		g.Iterate()
	}
	took := time.Since(start_time)
	fmt.Printf("engine: %s\n%dx%d board\n%d iterations\ntook %s\n", *engine_name, board.Rows, board.Cols, *iters, took)
	return took, nil
}

func bench(board boards.BoardSpec) (time.Duration, error) {
	total_dur_sec := 0.0

	for range *cnt {
		dur, err := run(board)
		if err != nil {
			log.Fatalf("err: %v", err)
		}
		total_dur_sec += dur.Seconds()
	}
	total_dur_sec /= float64(*cnt)

	avg_dur := time.Duration(total_dur_sec * float64(time.Second))
	return avg_dur, nil
}

func main() {
	flag.Parse()
	board := boards.Random(*rows, *cols, *seed)

	avg_dur, err := bench(board)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("engine %s\navg dur: %s\n", *engine_name, avg_dur)
}
