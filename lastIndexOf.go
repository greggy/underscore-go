package underscorego

func LastIndexOf[T comparable](x []T, value T) int {
	for i := len(x) - 1; i >= 0; i-- {
		if x[i] == value {
			return i
		}
	}
	return -1
}
