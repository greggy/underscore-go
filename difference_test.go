package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Difference_Property_NoneInExcluded(t *testing.T) {
	f := func(a, b []int8) bool {
		r := Difference(a, b)
		set := make(map[int8]bool)
		for _, e := range b {
			set[e] = true
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

func Test_Difference_Property_AllFromSource(t *testing.T) {
	f := func(a, b []int8) bool {
		r := Difference(a, b)
		set := make(map[int8]bool)
		for _, e := range a {
			set[e] = true
		}
		for _, e := range r {
			if !set[e] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Difference_Property_EmptyExclude(t *testing.T) {
	f := func(x []int) bool {
		r := Difference(x)
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

func Test_Difference_Property_ConsistentWithWithout(t *testing.T) {
	f := func(a, b []int8) bool {
		d := Difference(a, b)
		w := Without(a, b...)
		if len(d) != len(w) {
			return false
		}
		for i := range d {
			if d[i] != w[i] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}
