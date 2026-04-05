package underscorego

// First returns a new slice containing the first n elements of x.
// If n exceeds len(x), all elements are returned.
func First[T any](x []T, n int) []T {
	if n <= 0 {
		return []T{}
	}
	if n >= len(x) {
		n = len(x)
	}
	r := make([]T, n)
	copy(r, x[:n])
	return r
}
