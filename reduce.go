package underscorego

// Reduce reduces the slice x to a single value by applying fn from left to right,
// starting with the initial value m.
func Reduce[T any](x []T, fn func(T, T) T, m T) T {
	for _, e := range seqAll(x) {
		m = fn(m, e)
	}
	return m
}
