package underscorego

// Map returns a new slice with each element of x transformed by fn.
func Map[T, M any](x []T, fn func(T) M) []M {
	y := make([]M, 0, len(x))
	for _, e := range x {
		y = append(y, fn(e))
	}
	return y
}
