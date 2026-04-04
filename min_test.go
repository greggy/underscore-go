package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Min_Property_NoElementSmaller(t *testing.T) {
	f := func(x []int) bool {
		if len(x) == 0 {
			return true
		}
		m := Min(x)
		for _, v := range x {
			if v < m {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Min_Property_ElementExists(t *testing.T) {
	f := func(x []int) bool {
		if len(x) == 0 {
			return true
		}
		m := Min(x)
		return Contains(x, m)
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Min_Property_EmptyReturnsZero(t *testing.T) {
	require.Equal(t, 0, Min([]int{}))
	require.Equal(t, 0.0, Min([]float64{}))
}
