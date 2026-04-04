package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_FindIndex_Property_ResultSatisfies(t *testing.T) {
	f := func(x []int) bool {
		pred := func(i int) bool { return i > 50 }
		idx := FindIndex(x, pred)
		if idx == -1 {
			return !Some(x, pred)
		}
		return pred(x[idx])
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_FindIndex_Property_FirstMatch(t *testing.T) {
	f := func(x []int) bool {
		pred := func(i int) bool { return i%7 == 0 && i != 0 }
		idx := FindIndex(x, pred)
		if idx == -1 {
			return true
		}
		for i := 0; i < idx; i++ {
			if pred(x[i]) {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_FindIndex_Property_ConsistentWithFind(t *testing.T) {
	f := func(x []int) bool {
		pred := func(i int) bool { return i < -10 }
		idx := FindIndex(x, pred)
		found := Find(x, pred)
		if idx == -1 {
			return found == nil
		}
		return found != nil && *found == x[idx]
	}
	require.NoError(t, quick.Check(f, nil))
}
