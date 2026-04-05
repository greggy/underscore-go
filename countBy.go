package underscorego

// CountBy counts the elements of x grouped by the result of fn.
func CountBy[T any, K comparable](x []T, fn func(T) K) map[K]int {
	r := make(map[K]int)
	for _, e := range x {
		r[fn(e)]++
	}
	return r
}
