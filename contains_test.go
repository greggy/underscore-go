package underscorego

import (
	"sort"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Contains_Property_FoundElement(t *testing.T) {
	f := func(x []int, idx uint) bool {
		if len(x) == 0 {
			return true
		}
		i := idx % uint(len(x))
		return Contains(x, x[i])
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Contains_Property_NotFoundElement(t *testing.T) {
	f := func(x []byte) bool {
		m := make(map[byte]bool)
		for _, e := range x {
			m[e] = true
		}
		for v := 0; v < 256; v++ {
			b := byte(v)
			if Contains(x, b) != m[b] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Contains_Property_EmptySlice(t *testing.T) {
	f := func(v int) bool {
		return !Contains([]int{}, v)
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Contains_Property_ConsistentWithIndexOf(t *testing.T) {
	f := func(x []int, v int) bool {
		return Contains(x, v) == (IndexOf(x, v) >= 0)
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Contains_Property_SortedBinaryConsistency(t *testing.T) {
	f := func(x []int, v int) bool {
		sort.Ints(x)
		idx := sort.SearchInts(x, v)
		found := idx < len(x) && x[idx] == v
		return Contains(x, v) == found
	}
	require.NoError(t, quick.Check(f, nil))
}
