package underscorego

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
