package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Reduce_Property_SumEqualsLoop(t *testing.T) {
	f := func(x []int) bool {
		sum := 0
		for _, e := range x {
			sum += e
		}
		return Reduce(x, func(a, b int) int { return a + b }, 0) == sum
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Reduce_Property_ConcatPreservesAll(t *testing.T) {
	f := func(x []string) bool {
		concat := Reduce(x, func(a, b string) string { return a + b }, "")
		total := 0
		for _, s := range x {
			total += len(s)
		}
		return len(concat) == total
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Reduce_Property_EmptyReturnsMemo(t *testing.T) {
	f := func(m int) bool {
		return Reduce([]int{}, func(a, b int) int { return a + b }, m) == m
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Reduce_Property_SingleElement(t *testing.T) {
	f := func(v int, m int) bool {
		return Reduce([]int{v}, func(a, b int) int { return a + b }, m) == m+v
	}
	require.NoError(t, quick.Check(f, nil))
}
