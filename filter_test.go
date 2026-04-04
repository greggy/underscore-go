package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Filter_Property_AllMatchPredicate(t *testing.T) {
	f := func(x []int) bool {
		pred := func(i int) bool { return i%2 == 0 }
		for _, e := range Filter(x, pred) {
			if !pred(e) {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Filter_Property_ComplementWithReject(t *testing.T) {
	f := func(x []int) bool {
		pred := func(i int) bool { return i > 0 }
		filtered := Filter(x, pred)
		rejected := Reject(x, pred)
		return len(filtered)+len(rejected) == len(x)
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Filter_Property_PreservesOrder(t *testing.T) {
	f := func(x []int) bool {
		pred := func(i int) bool { return i%3 == 0 }
		r := Filter(x, pred)
		j := 0
		for _, e := range x {
			if pred(e) {
				if r[j] != e {
					return false
				}
				j++
			}
		}
		return j == len(r)
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Filter_Property_TruePredicateIsIdentity(t *testing.T) {
	f := func(x []int) bool {
		r := Filter(x, func(int) bool { return true })
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

func Test_Filter_Property_FalsePredicateIsEmpty(t *testing.T) {
	f := func(x []int) bool {
		r := Filter(x, func(int) bool { return false })
		return len(r) == 0
	}
	require.NoError(t, quick.Check(f, nil))
}
