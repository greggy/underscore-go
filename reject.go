package underscorego

func Reject[T any](x []T, fn func(T) bool) []T {
	r := []T{}
	for _, e := range x {
		if !fn(e) {
			r = append(r, e)
		}
	}
	return r
}
