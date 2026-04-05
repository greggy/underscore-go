package underscorego

import "math/rand"

// Sample returns n random unique elements from x.
// If n >= len(x), all elements are returned in random order.
func Sample[T any](x []T, n int) []T {
	if n <= 0 || len(x) == 0 {
		return []T{}
	}
	if n >= len(x) {
		n = len(x)
	}
	shuffled := make([]T, len(x))
	copy(shuffled, x)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled[:n]
}
