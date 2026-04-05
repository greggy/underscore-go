package underscorego

// Range generates a slice of integers. It accepts one to three arguments:
//   - Range(stop): generates [0, 1, ..., stop-1]
//   - Range(start, stop): generates [start, start+1, ..., stop-1]
//   - Range(start, stop, step): generates values from start to stop (exclusive) with the given step
func Range(args ...int) []int {
	var start, stop, step int
	switch len(args) {
	case 1:
		stop = args[0]
		if stop > 0 {
			step = 1
		} else {
			step = -1
		}
	case 2:
		start = args[0]
		stop = args[1]
		if start <= stop {
			step = 1
		} else {
			step = -1
		}
	case 3:
		start = args[0]
		stop = args[1]
		step = args[2]
	default:
		return []int{}
	}
	if step == 0 {
		return []int{}
	}
	r := []int{}
	if step > 0 {
		for i := start; i < stop; i += step {
			r = append(r, i)
		}
	} else {
		for i := start; i > stop; i += step {
			r = append(r, i)
		}
	}
	return r
}
