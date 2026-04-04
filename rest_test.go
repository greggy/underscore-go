package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Rest_Property_RemovesFirstN(t *testing.T) {
	f := func(x []int) bool {
		if len(x) < 2 {
			return true
		}
		r := Rest(x, 1)
		return len(r) == len(x)-1
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Rest_Property_MatchesSuffix(t *testing.T) {
	f := func(x []int) bool {
		if len(x) == 0 {
			return true
		}
		n := len(x) / 2
		r := Rest(x, n)
		for i := range r {
			if r[i] != x[n+i] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Rest_Property_ZeroReturnsAll(t *testing.T) {
	f := func(x []int) bool {
		r := Rest(x, 0)
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
