package underscorego

import "math/rand"

// Shuffle returns a randomly shuffled copy of x.
// The original slice is not modified.
func Shuffle[T any](x []T) []T {
	y := make([]T, len(x))
	copy(y, x)
	rand.Shuffle(len(y), func(i, j int) {
		y[i], y[j] = y[j], y[i]
	})
	return y
}
