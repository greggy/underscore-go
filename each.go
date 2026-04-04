package underscorego

func Each[T any](x []T, fn func(T)) {
	for _, e := range x {
		fn(e)
	}
}
