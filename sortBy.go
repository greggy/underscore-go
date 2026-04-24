package underscorego

import (
	"cmp"
	"sort"
)

// SortBy returns a sorted copy of x, ordered by the result of applying fn to each element.
// The original slice is not modified.
func SortBy[T cmp.Ordered](x []T, fn func(T) T) []T {
	type pair struct {
		val T
		key T
	}
	pairs := make([]pair, len(x))
	for i, e := range seqAll(x) {
		pairs[i] = pair{val: e, key: fn(e)}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].key < pairs[j].key
	})
	result := make([]T, len(x))
	for i, p := range seqAll(pairs) {
		result[i] = p.val
	}
	return result
}
