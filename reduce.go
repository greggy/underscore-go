package underscorego

func Reduce[T any](x []T, fn func(T, T) T, m T) T {
	for _, e := range x {
		m = fn(m, e)
	}
	return m
}
