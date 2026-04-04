package underscorego

func Find[T any](x []T, fn func(T) bool) *T {
	for _, e := range x {
		if fn(e) {
			result := e
			return &result
		}
	}
	return nil
}
