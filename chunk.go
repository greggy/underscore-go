package underscorego

// Chunk splits the slice x into groups of the given size.
// The last chunk may contain fewer elements.
func Chunk[T any](x []T, size int) [][]T {
	if size <= 0 || len(x) == 0 {
		return [][]T{}
	}
	r := [][]T{}
	for i := range seqRange(0, len(x), size) {
		end := i + size
		if end > len(x) {
			end = len(x)
		}
		chunk := make([]T, end-i)
		copy(chunk, x[i:end])
		r = append(r, chunk)
	}
	return r
}
