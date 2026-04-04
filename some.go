package underscorego

func Some[T any](x []T, fn func(T) bool) bool {
	for _, e := range x {
		if fn(e) {
			return true
		}
	}
	return false
}
