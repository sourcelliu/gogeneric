package slices

import "golang.org/x/exp/constraints"

func SliceToMap[T any, E constraints.Ordered](v []T, f func(T) E) (m map[E]T) {
	m = make(map[E]T, 0)
	for _, val := range v {
		m[f(val)] = val
	}
	return
}
