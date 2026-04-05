package underscorego

// Find returns a pointer to the first element in x for which fn returns true.
// Returns nil if no element matches.
func Find[T any](x []T, fn func(T) bool) *T {
	for _, e := range x {
		if fn(e) {
			result := e
			return &result
		}
	}
	return nil
}
