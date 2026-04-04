package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Intersection_Property_AllInBothArrays(t *testing.T) {
	f := func(a, b []int8) bool {
		r := Intersection(a, b)
		setA := make(map[int8]bool)
		for _, e := range a {
			setA[e] = true
		}
		setB := make(map[int8]bool)
		for _, e := range b {
			setB[e] = true
		}
		for _, e := range r {
			if !setA[e] || !setB[e] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Intersection_Property_SubsetOfBoth(t *testing.T) {
	f := func(a, b []int8) bool {
		r := Intersection(a, b)
		return len(r) <= len(a) && len(r) <= len(b)
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Intersection_Property_NoDuplicates(t *testing.T) {
	f := func(a, b []int8) bool {
		r := Intersection(a, b)
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

func Test_Intersection_Property_SelfEqualsUniq(t *testing.T) {
	f := func(x []int8) bool {
		r := Intersection(x, x)
		u := Uniq(x)
		if len(r) != len(u) {
			return false
		}
		set := make(map[int8]bool)
		for _, e := range r {
			set[e] = true
		}
		for _, e := range u {
			if !set[e] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}
