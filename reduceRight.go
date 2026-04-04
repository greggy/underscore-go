package underscorego

func ReduceRight[T any](x []T, fn func(T, T) T, m T) T {
	for i := len(x) - 1; i >= 0; i-- {
		m = fn(m, x[i])
	}
	return m
}
