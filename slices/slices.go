package slices

import "golang.org/x/exp/constraints"

func SliceToMap[E constraints.Ordered](v []E) (m map[E]bool) {
	if len(v) == 0 {
		return nil
	}
	m = make(map[E]bool, len(v))
	for _, val := range v {
		m[val] = true
	}
	return
}

func SliceToMapFunc[T any, E constraints.Ordered](v []T, f func(T) E) (m map[E]T) {
	if len(v) == 0 || f == nil {
		return nil
	}
	m = make(map[E]T, len(v))
	for _, val := range v {
		m[f(val)] = val
	}
	return
}
