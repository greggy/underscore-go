package underscorego

func Flatten[T any](x [][]T) []T {
	r := []T{}
	for _, inner := range x {
		r = append(r, inner...)
	}
	return r
}
