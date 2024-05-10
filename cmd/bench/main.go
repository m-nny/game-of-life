package main

import (
	"flag"
	"log"
	"time"

	"minmax.uk/game-of-life/pkg/engine"
)

var (
	rows  = flag.Int64("rows", 1000, "rows")
	cols  = flag.Int64("cols", 1000, "cols")
	seed  = flag.Int64("seed", 42, "seed")
	iters = flag.Int("iters", 1000, "iters")
	cnt   = flag.Int("cnt", 5, "number of times to run")
)

func bench(b engine.BoardSpec, iters int) (time.Duration, error) {
	g, err := engine.FromBoardSpec(b)
	if err != nil {
		return 0, err
	}
	start_time := time.Now()
	for range iters {
		g.Iterate()
	}
	took := time.Since(start_time)
	log.Printf("%dx%d board for %d iterations took %s", b.Rows, b.Cols, iters, took)
	return took, nil
}

func main() {
	flag.Parse()
	b := engine.Random(*rows, *cols, *seed)

	total_dur_sec := 0.0

	for range *cnt {
		dur, err := bench(b, *iters)
		if err != nil {
			log.Fatalf("err: %v", err)
		}
		total_dur_sec += dur.Seconds()
	}
	total_dur_sec /= float64(*cnt)

	avg_dur := time.Duration(total_dur_sec * float64(time.Second))

	log.Printf("avg dur: %s", avg_dur)
}
