package underscorego

import (
	"sort"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Sample_Property_CorrectSize(t *testing.T) {
	f := func(x []int, n uint8) bool {
		s := int(n) % 20
		r := Sample(x, s)
		expected := s
		if expected > len(x) {
			expected = len(x)
		}
		if s <= 0 {
			expected = 0
		}
		return len(r) == expected
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Sample_Property_ElementsFromSource(t *testing.T) {
	f := func(x []int) bool {
		if len(x) == 0 {
			return true
		}
		r := Sample(x, 3)
		set := make(map[int]bool)
		for _, e := range x {
			set[e] = true
		}
		for _, e := range r {
			if !set[e] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Sample_Property_NoDuplicates(t *testing.T) {
	f := func(x []uint8) bool {
		uniqX := Uniq(x)
		r := Sample(uniqX, len(uniqX))
		sorted := make([]int, len(r))
		for i, v := range r {
			sorted[i] = int(v)
		}
		sort.Ints(sorted)
		for i := 1; i < len(sorted); i++ {
			if sorted[i] == sorted[i-1] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}
