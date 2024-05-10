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

func (s HashSet[V]) Get(v V) (V, bool) {
	arr, ok := s[v.Hash()]
	if !ok {
		return v, false
	}
	idx := slices.IndexFunc(arr, func(item V) bool { return item.Equals(v) })
	if idx == -1 {
		return v, false
	}
	return arr[idx], true
}

func (s HashSet[V]) Contains(v V) bool {
	_, ok := s.Get(v)
	return ok
}

func (s HashSet[V]) Add(v V) {
	h := v.Hash()
	arr := s[h]
	if slices.ContainsFunc(arr, func(item V) bool { return item.Equals(v) }) {
		return
	}
	s[h] = append(arr, v)
}
