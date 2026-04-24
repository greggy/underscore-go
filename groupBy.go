package underscorego

// GroupBy groups the elements of x by the result of fn,
// returning a map from keys to slices of matching elements.
func GroupBy[T any, K comparable](x []T, fn func(T) K) map[K][]T {
	r := make(map[K][]T)
	for _, e := range seqAll(x) {
		k := fn(e)
		r[k] = append(r[k], e)
	}
	return r
}
