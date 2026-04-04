package underscorego

func Uniq[T comparable](x []T) []T {
	seen := make(map[T]bool)
	r := []T{}
	for _, e := range x {
		if !seen[e] {
			seen[e] = true
			r = append(r, e)
		}
	}
	return r
}
