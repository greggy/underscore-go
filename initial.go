package underscorego

// Initial returns all elements of x except the last n.
// If n >= len(x), returns an empty slice.
func Initial[T any](x []T, n int) []T {
	if n <= 0 || len(x) == 0 {
		r := make([]T, len(x))
		copy(r, x)
		return r
	}
	if n >= len(x) {
		return []T{}
	}
	r := make([]T, len(x)-n)
	copy(r, x[:len(x)-n])
	return r
}
