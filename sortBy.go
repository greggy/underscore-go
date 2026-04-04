package underscorego

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func SortBy[T constraints.Ordered](x []T, fn func(T) T) []T {
	type pair struct {
		val T
		key T
	}
	pairs := make([]pair, len(x))
	for i, e := range x {
		pairs[i] = pair{val: e, key: fn(e)}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].key < pairs[j].key
	})
	result := make([]T, len(x))
	for i, p := range pairs {
		result[i] = p.val
	}
	return result
}
