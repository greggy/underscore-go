package underscorego

import (
	"iter"
	"slices"
)

func seqValues[T any](x []T) iter.Seq[T] {
	return slices.Values(x)
}

func seqAll[T any](x []T) iter.Seq2[int, T] {
	return slices.All(x)
}

func seqBackward[T any](x []T) iter.Seq2[int, T] {
	return slices.Backward(x)
}

func seqRange(start, stop, step int) iter.Seq[int] {
	return func(yield func(int) bool) {
		if step == 0 {
			return
		}
		if step > 0 {
			for i := start; i < stop; i += step {
				if !yield(i) {
					return
				}
			}
			return
		}
		for i := start; i > stop; i += step {
			if !yield(i) {
				return
			}
		}
	}
}
