package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Reject_Property_NoneMatchPredicate(t *testing.T) {
	f := func(x []int) bool {
		pred := func(i int) bool { return i%2 == 0 }
		for _, e := range Reject(x, pred) {
			if pred(e) {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Reject_Property_InverseOfFilter(t *testing.T) {
	f := func(x []int) bool {
		pred := func(i int) bool { return i > 0 }
		filtered := Filter(x, pred)
		rejected := Reject(x, pred)
		if len(filtered)+len(rejected) != len(x) {
			return false
		}
		j, k := 0, 0
		for _, e := range x {
			if pred(e) {
				if j >= len(filtered) || filtered[j] != e {
					return false
				}
				j++
			} else {
				if k >= len(rejected) || rejected[k] != e {
					return false
				}
				k++
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Reject_Property_FalsePredicateIsIdentity(t *testing.T) {
	f := func(x []int) bool {
		r := Reject(x, func(int) bool { return false })
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

func Test_Reject_Property_TruePredicateIsEmpty(t *testing.T) {
	f := func(x []int) bool {
		return len(Reject(x, func(int) bool { return true })) == 0
	}
	require.NoError(t, quick.Check(f, nil))
}
