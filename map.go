package underscorego

func Map[T, M any](x []T, fn func(T) M) []M {
	y := []M{}
	for _, e := range x {
		y = append(y, fn(e))
	}
	return y
}
