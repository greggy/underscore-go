package underscorego

// LastIndexOf returns the index of the last occurrence of value in x, or -1 if not found.
func LastIndexOf[T comparable](x []T, value T) int {
	for i, e := range seqBackward(x) {
		if e == value {
			return i
		}
	}
	return -1
}
