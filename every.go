package underscorego

func Every[T any](x []T, fn func(T) bool) bool {
	if len(x) == 0 {
		return false
	}
	for _, e := range x {
		if !fn(e) {
			return false
		}
	}
	return true
}
