package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Find_Property_ResultSatisfiesPredicate(t *testing.T) {
	f := func(x []int) bool {
		pred := func(i int) bool { return i%5 == 0 }
		r := Find(x, pred)
		if r != nil && !pred(*r) {
			return false
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Find_Property_ReturnsFirstMatch(t *testing.T) {
	f := func(x []int) bool {
		pred := func(i int) bool { return i > 0 }
		r := Find(x, pred)
		if r == nil {
			return !Some(x, pred)
		}
		idx := FindIndex(x, pred)
		return x[idx] == *r
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Find_Property_NilWhenNoneMatch(t *testing.T) {
	f := func(x []int) bool {
		pred := func(int) bool { return false }
		return Find(x, pred) == nil
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Find_Property_ConsistentWithSome(t *testing.T) {
	f := func(x []int) bool {
		pred := func(i int) bool { return i%3 == 0 && i != 0 }
		found := Find(x, pred)
		hasSome := Some(x, pred)
		return (found != nil) == hasSome
	}
	require.NoError(t, quick.Check(f, nil))
}
