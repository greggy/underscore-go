package underscorego

import (
	"sort"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Shuffle_Property_PreservesLength(t *testing.T) {
	f := func(x []int) bool {
		return len(Shuffle(x)) == len(x)
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Shuffle_Property_PreservesElements(t *testing.T) {
	f := func(x []int) bool {
		r := Shuffle(x)
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

func Test_Shuffle_Property_DoesNotMutateInput(t *testing.T) {
	f := func(x []int) bool {
		original := make([]int, len(x))
		copy(original, x)
		Shuffle(x)
		for i := range x {
			if x[i] != original[i] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}
