package slices

import "golang.org/x/exp/constraints"

func SliceToMap[T any, E constraints.Ordered](v []T, f func(T) E) (m map[E]T) {
	if len(v) == 0 || f == nil {
		return nil
	}
	m = make(map[E]T, 0)
	for _, val := range v {
		m[f(val)] = val
	}
	return
}
