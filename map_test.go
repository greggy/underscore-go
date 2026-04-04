package underscorego

import (
	"strconv"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Map_Property_PreservesLength(t *testing.T) {
	f := func(x []int) bool {
		r := Map(x, func(i int) int { return i * 2 })
		return len(r) == len(x)
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Map_Property_TransformApplied(t *testing.T) {
	f := func(x []int) bool {
		fn := func(i int) int { return i * 3 }
		r := Map(x, fn)
		for i, e := range x {
			if r[i] != fn(e) {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Map_Property_IdentityIsNoop(t *testing.T) {
	f := func(x []int) bool {
		r := Map(x, func(i int) int { return i })
		for i := range x {
			if r[i] != x[i] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Map_Property_CompositionEquivalence(t *testing.T) {
	f := func(x []int) bool {
		double := func(i int) int { return i * 2 }
		toString := strconv.Itoa
		composed := func(i int) string { return toString(double(i)) }

		r1 := Map(Map(x, double), toString)
		r2 := Map(x, composed)
		if len(r1) != len(r2) {
			return false
		}
		for i := range r1 {
			if r1[i] != r2[i] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}
