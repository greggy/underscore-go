package underscorego

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
