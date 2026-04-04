package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Each_Property_VisitsAllElements(t *testing.T) {
	f := func(x []int) bool {
		visited := []int{}
		Each(x, func(e int) { visited = append(visited, e) })
		if len(visited) != len(x) {
			return false
		}
		for i := range x {
			if visited[i] != x[i] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Each_Property_OrderPreserved(t *testing.T) {
	f := func(x []int) bool {
		idx := 0
		Each(x, func(e int) {
			if e != x[idx] {
				panic("order mismatch")
			}
			idx++
		})
		return idx == len(x)
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Each_Property_CountMatchesLength(t *testing.T) {
	f := func(x []string) bool {
		count := 0
		Each(x, func(_ string) { count++ })
		return count == len(x)
	}
	require.NoError(t, quick.Check(f, nil))
}
