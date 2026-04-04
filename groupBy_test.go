package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_GroupBy_Property_AllElementsPresent(t *testing.T) {
	f := func(x []int) bool {
		fn := func(i int) int { return i % 3 }
		groups := GroupBy(x, fn)
		total := 0
		for _, v := range groups {
			total += len(v)
		}
		return total == len(x)
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_GroupBy_Property_AllInGroupMatchKey(t *testing.T) {
	f := func(x []int) bool {
		fn := func(i int) int { return i % 5 }
		for k, v := range GroupBy(x, fn) {
			for _, e := range v {
				if fn(e) != k {
					return false
				}
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_GroupBy_Property_ConsistentWithCountBy(t *testing.T) {
	f := func(x []int) bool {
		fn := func(i int) int { return i % 4 }
		groups := GroupBy(x, fn)
		counts := CountBy(x, fn)
		for k, v := range groups {
			if counts[k] != len(v) {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}
