package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Every_Property_EmptyIsFalse(t *testing.T) {
	y := Every([]int{}, func(int) bool { return true })
	require.False(t, y)
}

func Test_Every_Property_NegationOfSome(t *testing.T) {
	f := func(x []int) bool {
		if len(x) == 0 {
			return true
		}
		pred := func(i int) bool { return i%2 == 0 }
		negPred := func(i int) bool { return !pred(i) }
		return Every(x, pred) == !Some(x, negPred)
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Every_Property_AllTrueIsTrue(t *testing.T) {
	f := func(x []int) bool {
		if len(x) == 0 {
			return true
		}
		return Every(x, func(int) bool { return true })
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Every_Property_OneFalseIsFalse(t *testing.T) {
	f := func(x []int) bool {
		if len(x) == 0 {
			return true
		}
		return !Every(x, func(int) bool { return false })
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Every_Property_SubsetImplication(t *testing.T) {
	f := func(x []int) bool {
		if len(x) == 0 {
			return true
		}
		strict := func(i int) bool { return i%6 == 0 }
		loose := func(i int) bool { return i%2 == 0 }
		if Every(x, strict) && !Every(x, loose) {
			return false
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}
