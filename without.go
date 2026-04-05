package underscorego

// Without returns a copy of x with all occurrences of the specified values removed.
func Without[T comparable](x []T, values ...T) []T {
	exclude := make(map[T]bool, len(values))
	for _, v := range values {
		exclude[v] = true
	}
	r := []T{}
	for _, e := range x {
		if !exclude[e] {
			r = append(r, e)
		}
	}
	return r
}
