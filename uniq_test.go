package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Uniq_Property_NoDuplicates(t *testing.T) {
	f := func(x []int8) bool {
		r := Uniq(x)
		seen := make(map[int8]bool)
		for _, e := range r {
			if seen[e] {
				return false
			}
			seen[e] = true
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Uniq_Property_PreservesAllUniqueValues(t *testing.T) {
	f := func(x []int8) bool {
		r := Uniq(x)
		set := make(map[int8]bool)
		for _, e := range r {
			set[e] = true
		}
		for _, e := range x {
			if !set[e] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Uniq_Property_Idempotent(t *testing.T) {
	f := func(x []int8) bool {
		r1 := Uniq(x)
		r2 := Uniq(r1)
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

func Test_Uniq_Property_FirstOccurrencePreserved(t *testing.T) {
	f := func(x []int8) bool {
		r := Uniq(x)
		for _, e := range r {
			first := IndexOf(x, e)
			if x[first] != e {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}
