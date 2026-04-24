package underscorego

// Difference returns the elements from x that are not present in any of the other slices.
func Difference[T comparable](x []T, others ...[]T) []T {
	exclude := make(map[T]bool)
	for arr := range seqValues(others) {
		for _, e := range seqAll(arr) {
			exclude[e] = true
		}
	}
	r := []T{}
	for _, e := range seqAll(x) {
		if !exclude[e] {
			r = append(r, e)
		}
	}
	return r
}
