package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_FindLastIndex_Property_ResultSatisfies(t *testing.T) {
	f := func(x []int) bool {
		pred := func(i int) bool { return i > 50 }
		idx := FindLastIndex(x, pred)
		if idx == -1 {
			return !Some(x, pred)
		}
		return pred(x[idx])
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_FindLastIndex_Property_LastMatch(t *testing.T) {
	f := func(x []int) bool {
		pred := func(i int) bool { return i%7 == 0 && i != 0 }
		idx := FindLastIndex(x, pred)
		if idx == -1 {
			return true
		}
		for i := idx + 1; i < len(x); i++ {
			if pred(x[i]) {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_FindLastIndex_Property_GeqFindIndex(t *testing.T) {
	f := func(x []int) bool {
		pred := func(i int) bool { return i%3 == 0 && i != 0 }
		return FindLastIndex(x, pred) >= FindIndex(x, pred)
	}
	require.NoError(t, quick.Check(f, nil))
}
