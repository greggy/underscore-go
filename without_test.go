package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Without_Property_ExcludedNotPresent(t *testing.T) {
	f := func(x []int8, vals []int8) bool {
		r := Without(x, vals...)
		set := make(map[int8]bool)
		for _, v := range vals {
			set[v] = true
		}
		for _, e := range r {
			if set[e] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Without_Property_PreservesOrder(t *testing.T) {
	f := func(x []int) bool {
		r := Without(x, -999)
		j := 0
		for _, e := range x {
			if e != -999 {
				if r[j] != e {
					return false
				}
				j++
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Without_Property_EmptyValsIsIdentity(t *testing.T) {
	f := func(x []int) bool {
		r := Without(x)
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
