package underscorego

func Contains[T comparable](x []T, y T) bool {
	for _, e := range x {
		if e == y {
			return true
		}
	}
	return false
}
