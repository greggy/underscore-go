package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_First_Property_CorrectLength(t *testing.T) {
	f := func(x []int, n uint8) bool {
		k := int(n) % 20
		r := First(x, k)
		expected := k
		if expected > len(x) {
			expected = len(x)
		}
		if k <= 0 {
			expected = 0
		}
		return len(r) == expected
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_First_Property_MatchesPrefix(t *testing.T) {
	f := func(x []int) bool {
		if len(x) == 0 {
			return true
		}
		n := len(x) / 2
		r := First(x, n)
		for i := range r {
			if r[i] != x[i] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_First_Property_ConcatWithRest(t *testing.T) {
	f := func(x []int, n uint8) bool {
		k := int(n) % 20
		head := First(x, k)
		tail := Rest(x, k)
		if len(head)+len(tail) != len(x) {
			return false
		}
		combined := append(head, tail...)
		for i := range x {
			if combined[i] != x[i] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}
