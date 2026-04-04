package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Flatten_Property_TotalLength(t *testing.T) {
	f := func(a, b, c []int) bool {
		nested := [][]int{a, b, c}
		r := Flatten(nested)
		return len(r) == len(a)+len(b)+len(c)
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Flatten_Property_PreservesOrder(t *testing.T) {
	f := func(a, b []int) bool {
		r := Flatten([][]int{a, b})
		for i := range a {
			if r[i] != a[i] {
				return false
			}
		}
		for i := range b {
			if r[len(a)+i] != b[i] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Flatten_Property_ChunkRoundTrip(t *testing.T) {
	f := func(x []int) bool {
		if len(x) == 0 {
			return true
		}
		chunks := Chunk(x, 3)
		r := Flatten(chunks)
		if len(r) != len(x) {
			return false
		}
		for i := range x {
			if r[i] != x[i] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}
