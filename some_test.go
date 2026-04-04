package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Some_Property_EmptyIsFalse(t *testing.T) {
	y := Some([]int{}, func(int) bool { return true })
	require.False(t, y)
}

func Test_Some_Property_NegationOfEvery(t *testing.T) {
	f := func(x []int) bool {
		if len(x) == 0 {
			return true
		}
		pred := func(i int) bool { return i%2 == 0 }
		negPred := func(i int) bool { return !pred(i) }
		return Some(x, pred) == !Every(x, negPred)
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Some_Property_TrueIfContainsMatch(t *testing.T) {
	f := func(x []int) bool {
		pred := func(i int) bool { return i > 100 }
		hasSome := Some(x, pred)
		found := false
		for _, e := range x {
			if pred(e) {
				found = true
				break
			}
		}
		return hasSome == found
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Some_Property_ConsistentWithFind(t *testing.T) {
	f := func(x []int) bool {
		pred := func(i int) bool { return i < -50 }
		return Some(x, pred) == (Find(x, pred) != nil)
	}
	require.NoError(t, quick.Check(f, nil))
}
