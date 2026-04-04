package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_ReduceRight_Property_ReverseReduce(t *testing.T) {
	f := func(x []string) bool {
		fn := func(a, b string) string { return a + b }
		reversed := make([]string, len(x))
		for i, e := range x {
			reversed[len(x)-1-i] = e
		}
		return ReduceRight(x, fn, "") == Reduce(reversed, fn, "")
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_ReduceRight_Property_EmptyReturnsMemo(t *testing.T) {
	f := func(m int) bool {
		return ReduceRight([]int{}, func(a, b int) int { return a + b }, m) == m
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_ReduceRight_Property_CommutativeOpSameAsReduce(t *testing.T) {
	f := func(x []int) bool {
		add := func(a, b int) int { return a + b }
		return ReduceRight(x, add, 0) == Reduce(x, add, 0)
	}
	require.NoError(t, quick.Check(f, nil))
}
