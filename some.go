package underscorego

// Some returns true if at least one element in x satisfies the predicate fn.
func Some[T any](x []T, fn func(T) bool) bool {
	for _, e := range x {
		if fn(e) {
			return true
		}
	}
	return false
}
