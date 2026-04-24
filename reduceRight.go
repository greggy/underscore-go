package underscorego

// ReduceRight reduces the slice x to a single value by applying fn from right to left,
// starting with the initial value m.
func ReduceRight[T any](x []T, fn func(T, T) T, m T) T {
	for _, e := range seqBackward(x) {
		m = fn(m, e)
	}
	return m
}
