package underscorego

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func Test_Chunk_Property_FlattenRoundTrip(t *testing.T) {
	f := func(x []int) bool {
		if len(x) == 0 {
			return true
		}
		r := Flatten(Chunk(x, 3))
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

func Test_Chunk_Property_ChunkSizes(t *testing.T) {
	f := func(x []int) bool {
		if len(x) == 0 {
			return true
		}
		chunks := Chunk(x, 4)
		for i, c := range chunks {
			if i < len(chunks)-1 {
				if len(c) != 4 {
					return false
				}
			} else {
				if len(c) == 0 || len(c) > 4 {
					return false
				}
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Chunk_Property_ZeroSizeIsEmpty(t *testing.T) {
	f := func(x []int) bool {
		return len(Chunk(x, 0)) == 0
	}
	require.NoError(t, quick.Check(f, nil))
}

func Test_Chunk_Property_SizeOneIsSingleton(t *testing.T) {
	f := func(x []int) bool {
		chunks := Chunk(x, 1)
		if len(chunks) != len(x) {
			return false
		}
		for i, c := range chunks {
			if len(c) != 1 || c[0] != x[i] {
				return false
			}
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}
