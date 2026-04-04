package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_IndexOf_Property_FoundAtIndex(t *testing.T) {
	f := func(x []int, v int) bool {
		idx := IndexOf(x, v)
		if idx == -1 {
			return !Contains(x, v)
		}
		return x[idx] == v
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_IndexOf_Property_FirstOccurrence(t *testing.T) {
	f := func(x []int, v int) bool {
		idx := IndexOf(x, v)
		if idx == -1 {
			return true
		}
		for i := 0; i < idx; i++ {
			if x[i] == v {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_IndexOf_Property_ConsistentWithContains(t *testing.T) {
	f := func(x []int, v int) bool {
		return Contains(x, v) == (IndexOf(x, v) >= 0)
	}
	require.NoError(t, quick.Check(f, nil))
}
