package underscorego

import "cmp"

// Min returns the minimum element in x.
// Returns the zero value if x is empty.
func Min[T cmp.Ordered](x []T) T {
	if len(x) == 0 {
		var zero T
		return zero
	}
	m := x[0]
	for _, e := range seqAll(x[1:]) {
		if e < m {
			m = e
		}
	}
	return m
}
