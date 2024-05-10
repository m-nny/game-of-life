package hashset

import (
	"fmt"
	"slices"
)

type Hash = uint64

type Hashable interface {
	Hash() Hash
	Equals(other Hashable) bool
}

type HashSet[V Hashable] struct {
	m map[Hash][]V

	Stats struct {
		GetCnt        int
		GetSliceSizes int
		GetHit        int

		AddCnt        int
		AddSliceSizes int
		AddHit        int
	}
}

func New[V Hashable]() *HashSet[V] {
	m := make(map[Hash][]V)
	return &HashSet[V]{
		m: m,
	}
}

func (s *HashSet[V]) Get(v V) (V, bool) {
	s.Stats.GetCnt++
	arr, ok := s.m[v.Hash()]
	if !ok {
		return v, false
	}
	s.Stats.GetSliceSizes += len(arr)
	idx := slices.IndexFunc(arr, func(item V) bool { return item.Equals(v) })
	if idx == -1 {
		return v, false
	}
	s.Stats.GetHit++
	return arr[idx], true
}

func (s *HashSet[V]) Contains(v V) bool {
	_, ok := s.Get(v)
	return ok
}

func (s *HashSet[V]) Add(v V) {
	s.Stats.AddCnt++
	h := v.Hash()
	arr := s.m[h]
	s.Stats.AddSliceSizes += len(arr)
	// if slices.ContainsFunc(arr, func(item V) bool { return item.Equals(v) }) {
	// 	s.Stats.AddHit++
	// 	return
	// }
	s.m[h] = append(arr, v)
}

func (s *HashSet[V]) PrintStats() {
	fmt.Print("cell_cache_stats:\n")
	fmt.Print("  Get:\n")
	fmt.Printf("    Cnt: %d\n", s.Stats.GetCnt)
	fmt.Printf("    Hit: %d\n", s.Stats.GetHit)
	fmt.Printf("    AvgSliceSize: %.2f\n", float64(s.Stats.AddSliceSizes)/float64(s.Stats.GetCnt))
	fmt.Print("  Add:\n")
	fmt.Printf("    Cnt: %d\n", s.Stats.AddCnt)
	fmt.Printf("    Hit: %d\n", s.Stats.AddHit)
	fmt.Printf("    AvgSliceSize: %.2f\n", float64(s.Stats.AddSliceSizes)/float64(s.Stats.AddCnt))
}

func (s *HashSet[V]) ResetStats() {
	s.Stats.GetCnt = 0
	s.Stats.GetSliceSizes = 0
	s.Stats.AddCnt = 0
	s.Stats.AddSliceSizes = 0
}
