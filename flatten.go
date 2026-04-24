package underscorego

// Flatten flattens a slice of slices into a single slice.
func Flatten[T any](x [][]T) []T {
	n := 0
	for _, inner := range seqAll(x) {
		n += len(inner)
	}
	r := make([]T, 0, n)
	for _, inner := range seqAll(x) {
		r = append(r, inner...)
	}
	return r
}
