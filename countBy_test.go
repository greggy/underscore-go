package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_CountBy_Property_SumEqualsLength(t *testing.T) {
	f := func(x []int) bool {
		fn := func(i int) int { return i % 3 }
		counts := CountBy(x, fn)
		total := 0
		for _, c := range counts {
			total += c
		}
		return total == len(x)
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_CountBy_Property_AllCountsPositive(t *testing.T) {
	f := func(x []int) bool {
		fn := func(i int) bool { return i > 0 }
		for _, c := range CountBy(x, fn) {
			if c <= 0 {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_CountBy_Property_EmptySliceEmptyMap(t *testing.T) {
	counts := CountBy([]int{}, func(i int) int { return i })
	require.Empty(t, counts)
}
