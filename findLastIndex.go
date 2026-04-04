package underscorego

func FindLastIndex[T any](x []T, fn func(T) bool) int {
	for i := len(x) - 1; i >= 0; i-- {
		if fn(x[i]) {
			return i
		}
	}
	return -1
}
