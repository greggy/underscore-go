package underscorego

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Range_SingleArg(t *testing.T) {
	require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, Range(10))
}

func Test_Range_TwoArgs(t *testing.T) {
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, Range(1, 11))
}

func Test_Range_Step(t *testing.T) {
	require.Equal(t, []int{0, 5, 10, 15, 20, 25}, Range(0, 30, 5))
}

func Test_Range_NegativeStep(t *testing.T) {
	require.Equal(t, []int{0, -1, -2, -3, -4, -5, -6, -7, -8, -9}, Range(0, -10, -1))
}

func Test_Range_ZeroProducesEmpty(t *testing.T) {
	require.Empty(t, Range(0))
}

func Test_Range_Property_Length(t *testing.T) {
	r := Range(0, 100, 3)
	require.Equal(t, 34, len(r))
}

func Test_Range_Property_MonotonicallyIncreasing(t *testing.T) {
	r := Range(5, 50, 2)
	for i := 1; i < len(r); i++ {
		require.Greater(t, r[i], r[i-1])
	}
}

func Test_Range_Property_MonotonicallyDecreasing(t *testing.T) {
	r := Range(10, 0, -1)
	for i := 1; i < len(r); i++ {
		require.Less(t, r[i], r[i-1])
	}
}

func Test_Range_Property_StartEqualsFirst(t *testing.T) {
	r := Range(3, 10)
	require.Equal(t, 3, r[0])
}

func Test_Range_Property_StopExcluded(t *testing.T) {
	r := Range(0, 10)
	require.NotContains(t, r, 10)
}

func Test_Range_Property_DescendingAutoStep(t *testing.T) {
	r := Range(5, 0)
	require.Equal(t, []int{5, 4, 3, 2, 1}, r)
}
