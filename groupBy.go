package underscorego

func GroupBy[T any, K comparable](x []T, fn func(T) K) map[K][]T {
	r := make(map[K][]T)
	for _, e := range x {
		k := fn(e)
		r[k] = append(r[k], e)
	}
	return r
}
