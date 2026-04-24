package underscorego

// FindIndex returns the index of the first element in x for which fn returns true,
// or -1 if no element matches.
func FindIndex[T any](x []T, fn func(T) bool) int {
	for i, e := range seqAll(x) {
		if fn(e) {
			return i
		}
	}
	return -1
}
