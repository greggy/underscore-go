package underscorego

func FindIndex[T any](x []T, fn func(T) bool) int {
	for i, e := range x {
		if fn(e) {
			return i
		}
	}
	return -1
}
