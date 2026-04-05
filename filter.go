package underscorego

// Filter returns a new slice containing only the elements of x
// for which fn returns true.
func Filter[T any](x []T, fn func(T) bool) []T {
	r := []T{}
	for _, e := range x {
		if fn(e) {
			r = append(r, e)
		}
	}
	return r
}
