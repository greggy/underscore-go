package underscorego

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
