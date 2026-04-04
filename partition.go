package underscorego

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
