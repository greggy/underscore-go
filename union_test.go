package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Union_Property_NoDuplicates(t *testing.T) {
	f := func(a, b []int8) bool {
		r := Union(a, b)
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

func Test_Union_Property_ContainsAll(t *testing.T) {
	f := func(a, b []int8) bool {
		r := Union(a, b)
		set := make(map[int8]bool)
		for _, e := range r {
			set[e] = true
		}
		for _, e := range a {
			if !set[e] {
				return false
			}
		}
		for _, e := range b {
			if !set[e] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Union_Property_IdempotentWithSelf(t *testing.T) {
	f := func(x []int8) bool {
		r1 := Union(x)
		r2 := Union(r1, r1)
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
