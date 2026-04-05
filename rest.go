package underscorego

// Rest returns all elements of x except the first n.
// If n >= len(x), returns an empty slice.
func Rest[T any](x []T, n int) []T {
	if n <= 0 {
		r := make([]T, len(x))
		copy(r, x)
		return r
	}
	if n >= len(x) {
		return []T{}
	}
	r := make([]T, len(x)-n)
	copy(r, x[n:])
	return r
}
