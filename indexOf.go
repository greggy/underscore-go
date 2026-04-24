package underscorego

// IndexOf returns the index of the first occurrence of value in x, or -1 if not found.
func IndexOf[T comparable](x []T, value T) int {
	for i, e := range seqAll(x) {
		if e == value {
			return i
		}
	}
	return -1
}
