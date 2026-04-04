package underscorego

func IndexOf[T comparable](x []T, value T) int {
	for i, e := range x {
		if e == value {
			return i
		}
	}
	return -1
}
