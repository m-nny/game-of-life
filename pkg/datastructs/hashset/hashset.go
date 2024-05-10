package hashset

import "slices"

type Hash = uint64

type Hashable interface {
	Hash() Hash
	Equals(other Hashable) bool
}

type HashSet[V Hashable] map[Hash][]V

func New[V Hashable]() HashSet[V] {
	m := make(HashSet[V])
	return m
}

func (s HashSet[V]) Contains(v V) bool {
	arr, ok := s[v.Hash()]
	if !ok {
		return false
	}
	return slices.ContainsFunc(arr, func(item V) bool { return item.Equals(v) })
}

func (s HashSet[V]) Add(v V) {
	h := v.Hash()
	arr := s[h]
	if slices.ContainsFunc(arr, func(item V) bool { return item.Equals(v) }) {
		return
	}
	s[h] = append(arr, v)
}
