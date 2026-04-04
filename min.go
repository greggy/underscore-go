package underscorego

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](x []T) T {
	if len(x) == 0 {
		var zero T
		return zero
	}
	m := x[0]
	for _, e := range x[1:] {
		if e < m {
			m = e
		}
	}
	return m
}
