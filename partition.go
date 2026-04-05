package underscorego

// Partition splits x into two slices: the first contains elements for which fn
// returns true, the second contains the rest.
func Partition[T any](x []T, fn func(T) bool) ([]T, []T) {
	pass := []T{}
	fail := []T{}
	for _, e := range x {
		if fn(e) {
			pass = append(pass, e)
		} else {
			fail = append(fail, e)
		}
	}
	return pass, fail
}
