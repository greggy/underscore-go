package underscorego

// Intersection returns a slice of elements that are present in all provided slices.
func Intersection[T comparable](arrays ...[]T) []T {
	if len(arrays) == 0 {
		return []T{}
	}
	counts := make(map[T]int)
	for _, e := range seqAll(arrays[0]) {
		counts[e] = 1
	}
	for i := range seqRange(1, len(arrays), 1) {
		seen := make(map[T]bool)
		for _, e := range seqAll(arrays[i]) {
			if !seen[e] && counts[e] == i {
				counts[e]++
				seen[e] = true
			}
		}
	}
	r := []T{}
	seen := make(map[T]bool)
	for _, e := range seqAll(arrays[0]) {
		if !seen[e] && counts[e] == len(arrays) {
			r = append(r, e)
			seen[e] = true
		}
	}
	return r
}
