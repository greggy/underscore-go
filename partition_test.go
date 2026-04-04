package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Partition_Property_SumOfParts(t *testing.T) {
	f := func(x []int) bool {
		pass, fail := Partition(x, func(n int) bool { return n%2 == 0 })
		return len(pass)+len(fail) == len(x)
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Partition_Property_PassAllMatch(t *testing.T) {
	f := func(x []int) bool {
		pred := func(n int) bool { return n > 0 }
		pass, _ := Partition(x, pred)
		for _, v := range pass {
			if !pred(v) {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Partition_Property_FailNoneMatch(t *testing.T) {
	f := func(x []int) bool {
		pred := func(n int) bool { return n > 0 }
		_, fail := Partition(x, pred)
		for _, v := range fail {
			if pred(v) {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Partition_Property_EquivFilterReject(t *testing.T) {
	f := func(x []int) bool {
		pred := func(n int) bool { return n%3 == 0 }
		pass, fail := Partition(x, pred)
		filtered := Filter(x, pred)
		rejected := Reject(x, pred)
		if len(pass) != len(filtered) || len(fail) != len(rejected) {
			return false
		}
		for i := range pass {
			if pass[i] != filtered[i] {
				return false
			}
		}
		for i := range fail {
			if fail[i] != rejected[i] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}
