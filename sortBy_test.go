package underscorego

import (
	"sort"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_SortBy_Property_PreservesElements(t *testing.T) {
	f := func(x []int) bool {
		fn := func(i int) int { return i }
		r := SortBy(x, fn)
		if len(r) != len(x) {
			return false
		}
		sx := make([]int, len(x))
		copy(sx, x)
		sort.Ints(sx)
		sr := make([]int, len(r))
		copy(sr, r)
		sort.Ints(sr)
		for i := range sx {
			if sx[i] != sr[i] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_SortBy_Property_IdentityIsSorted(t *testing.T) {
	f := func(x []int) bool {
		r := SortBy(x, func(i int) int { return i })
		for i := 1; i < len(r); i++ {
			if r[i-1] > r[i] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_SortBy_Property_DoesNotMutateInput(t *testing.T) {
	f := func(x []int) bool {
		original := make([]int, len(x))
		copy(original, x)
		SortBy(x, func(i int) int { return -i })
		for i := range x {
			if x[i] != original[i] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}
