package underscorego

// Last returns a new slice containing the last n elements of x.
// If n exceeds len(x), all elements are returned.
func Last[T any](x []T, n int) []T {
	if n <= 0 {
		return []T{}
	}
	if n >= len(x) {
		n = len(x)
	}
	r := make([]T, n)
	copy(r, x[len(x)-n:])
	return r
}
