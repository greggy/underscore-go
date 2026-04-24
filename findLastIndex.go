package underscorego

// FindLastIndex returns the index of the last element in x for which fn returns true,
// or -1 if no element matches.
func FindLastIndex[T any](x []T, fn func(T) bool) int {
	for i, e := range seqBackward(x) {
		if fn(e) {
			return i
		}
	}
	return -1
}
