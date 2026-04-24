package underscorego

// Every returns true if all elements in x satisfy the predicate fn.
// Returns false for an empty slice.
func Every[T any](x []T, fn func(T) bool) bool {
	if len(x) == 0 {
		return false
	}
	for _, e := range seqAll(x) {
		if !fn(e) {
			return false
		}
	}
	return true
}
