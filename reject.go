package underscorego

// Reject returns a new slice containing the elements of x
// for which fn returns false.
func Reject[T any](x []T, fn func(T) bool) []T {
	r := []T{}
	for _, e := range seqAll(x) {
		if !fn(e) {
			r = append(r, e)
		}
	}
	return r
}
