package underscorego

func Union[T comparable](arrays ...[]T) []T {
	seen := make(map[T]bool)
	r := []T{}
	for _, arr := range arrays {
		for _, e := range arr {
			if !seen[e] {
				seen[e] = true
				r = append(r, e)
			}
		}
	}
	return r
}
