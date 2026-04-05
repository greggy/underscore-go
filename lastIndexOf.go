package underscorego

// LastIndexOf returns the index of the last occurrence of value in x, or -1 if not found.
func LastIndexOf[T comparable](x []T, value T) int {
	for i := len(x) - 1; i >= 0; i-- {
		if x[i] == value {
			return i
		}
	}
	return -1
}
