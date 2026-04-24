package underscorego

// Each iterates over the slice x, calling fn for each element.
func Each[T any](x []T, fn func(T)) {
	for _, e := range seqAll(x) {
		fn(e)
	}
}
