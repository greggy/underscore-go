package underscorego

// Union returns a slice of unique elements from all provided slices,
// preserving the order of first appearance.
func Union[T comparable](arrays ...[]T) []T {
	seen := make(map[T]bool)
	r := []T{}
	for arr := range seqValues(arrays) {
		for _, e := range seqAll(arr) {
			if !seen[e] {
				seen[e] = true
				r = append(r, e)
			}
		}
	}
	return r
}
