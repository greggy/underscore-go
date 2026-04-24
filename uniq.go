package underscorego

// Uniq returns a new slice with duplicate elements removed,
// preserving the order of first occurrence.
func Uniq[T comparable](x []T) []T {
	seen := make(map[T]bool)
	r := []T{}
	for _, e := range seqAll(x) {
		if !seen[e] {
			seen[e] = true
			r = append(r, e)
		}
	}
	return r
}
