package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Max_Property_NoElementLarger(t *testing.T) {
	f := func(x []int) bool {
		if len(x) == 0 {
			return true
		}
		m := Max(x)
		for _, e := range x {
			if e > m {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Max_Property_ElementExists(t *testing.T) {
	f := func(x []int) bool {
		if len(x) == 0 {
			return true
		}
		return Contains(x, Max(x))
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Max_Property_MinLeqMax(t *testing.T) {
	f := func(x []int) bool {
		if len(x) == 0 {
			return true
		}
		return Min(x) <= Max(x)
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Max_Property_EmptyReturnsZero(t *testing.T) {
	require.Equal(t, 0, Max([]int{}))
}
