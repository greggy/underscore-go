package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Last_Property_CorrectLength(t *testing.T) {
	f := func(x []int, n uint8) bool {
		k := int(n) % 20
		r := Last(x, k)
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

func Test_Last_Property_MatchesSuffix(t *testing.T) {
	f := func(x []int) bool {
		if len(x) == 0 {
			return true
		}
		n := len(x) / 2
		r := Last(x, n)
		offset := len(x) - n
		for i := range r {
			if r[i] != x[offset+i] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Last_Property_ConcatWithInitial(t *testing.T) {
	f := func(x []int, n uint8) bool {
		k := int(n) % 20
		head := Initial(x, k)
		tail := Last(x, k)
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
