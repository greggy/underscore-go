package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_LastIndexOf_Property_FoundAtIndex(t *testing.T) {
	f := func(x []int, v int) bool {
		idx := LastIndexOf(x, v)
		if idx == -1 {
			return !Contains(x, v)
		}
		return x[idx] == v
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_LastIndexOf_Property_LastOccurrence(t *testing.T) {
	f := func(x []int, v int) bool {
		idx := LastIndexOf(x, v)
		if idx == -1 {
			return true
		}
		for i := idx + 1; i < len(x); i++ {
			if x[i] == v {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_LastIndexOf_Property_GeqIndexOf(t *testing.T) {
	f := func(x []int, v int) bool {
		return LastIndexOf(x, v) >= IndexOf(x, v)
	}
	require.NoError(t, quick.Check(f, nil))
}
