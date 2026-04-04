package underscorego

func Flatten[T any](x [][]T) []T {
	n := 0
	for _, inner := range x {
		n += len(inner)
	}
	r := make([]T, 0, n)
	for _, inner := range x {
		r = append(r, inner...)
	}
	return r
}
