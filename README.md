# underscore-go

A Go generic utility library inspired by [Underscore.js](https://underscorejs.org). Provides functional helpers for slices, maps, and common collection operations using Go 1.18+ generics.

## Install

```bash
go get github.com/greggy/underscore-go
```

## Functions

### Iteration & Transformation

| Function | Signature | Description |
|----------|-----------|-------------|
| **Each** | `Each[T any](x []T, fn func(T))` | Iterates over a slice, calling `fn` for each element. |
| **Map** | `Map[T, M any](x []T, fn func(T) M) []M` | Returns a new slice with each element transformed by `fn`. |
| **Filter** | `Filter[T any](x []T, fn func(T) bool) []T` | Returns elements that satisfy the predicate. |
| **Reject** | `Reject[T any](x []T, fn func(T) bool) []T` | Returns elements that do not satisfy the predicate. |
| **Reduce** | `Reduce[T any](x []T, fn func(T, T) T, m T) T` | Reduces the slice to a single value from left to right. |
| **ReduceRight** | `ReduceRight[T any](x []T, fn func(T, T) T, m T) T` | Reduces the slice to a single value from right to left. |
| **Every** | `Every[T any](x []T, fn func(T) bool) bool` | Returns `true` if all elements satisfy the predicate. |
| **Some** | `Some[T any](x []T, fn func(T) bool) bool` | Returns `true` if at least one element satisfies the predicate. |
| **Partition** | `Partition[T any](x []T, fn func(T) bool) ([]T, []T)` | Splits the slice into two: elements that pass and elements that fail the predicate. |

### Search

| Function | Signature | Description |
|----------|-----------|-------------|
| **Contains** | `Contains[T comparable](x []T, y T) bool` | Returns `true` if the slice contains the value. |
| **Find** | `Find[T any](x []T, fn func(T) bool) *T` | Returns a pointer to the first element matching the predicate, or `nil`. |
| **FindIndex** | `FindIndex[T any](x []T, fn func(T) bool) int` | Returns the index of the first element matching the predicate, or `-1`. |
| **FindLastIndex** | `FindLastIndex[T any](x []T, fn func(T) bool) int` | Returns the index of the last element matching the predicate, or `-1`. |
| **IndexOf** | `IndexOf[T comparable](x []T, value T) int` | Returns the index of the first occurrence of the value, or `-1`. |
| **LastIndexOf** | `LastIndexOf[T comparable](x []T, value T) int` | Returns the index of the last occurrence of the value, or `-1`. |

### Slicing & Restructuring

| Function | Signature | Description |
|----------|-----------|-------------|
| **First** | `First[T any](x []T, n int) []T` | Returns the first `n` elements. |
| **Last** | `Last[T any](x []T, n int) []T` | Returns the last `n` elements. |
| **Initial** | `Initial[T any](x []T, n int) []T` | Returns all elements except the last `n`. |
| **Rest** | `Rest[T any](x []T, n int) []T` | Returns all elements except the first `n`. |
| **Flatten** | `Flatten[T any](x [][]T) []T` | Flattens a 2D slice into a 1D slice. |
| **Chunk** | `Chunk[T any](x []T, size int) [][]T` | Splits a slice into chunks of the given size. |

### Set Operations

| Function | Signature | Description |
|----------|-----------|-------------|
| **Uniq** | `Uniq[T comparable](x []T) []T` | Returns a slice with duplicate elements removed, preserving first occurrence order. |
| **Union** | `Union[T comparable](arrays ...[]T) []T` | Returns unique elements from all provided slices. |
| **Intersection** | `Intersection[T comparable](arrays ...[]T) []T` | Returns elements present in every provided slice. |
| **Difference** | `Difference[T comparable](x []T, others ...[]T) []T` | Returns elements from the first slice not present in the others. |
| **Without** | `Without[T comparable](x []T, values ...T) []T` | Returns a slice with the specified values removed. |

### Aggregation & Sorting

| Function | Signature | Description |
|----------|-----------|-------------|
| **Min** | `Min[T constraints.Ordered](x []T) T` | Returns the minimum element. Returns zero value for empty slices. |
| **Max** | `Max[T constraints.Ordered](x []T) T` | Returns the maximum element. Returns zero value for empty slices. |
| **SortBy** | `SortBy[T constraints.Ordered](x []T, fn func(T) T) []T` | Returns a sorted copy using `fn` as the sort key. Does not mutate input. |
| **GroupBy** | `GroupBy[T any, K comparable](x []T, fn func(T) K) map[K][]T` | Groups elements by the result of `fn`. |
| **CountBy** | `CountBy[T any, K comparable](x []T, fn func(T) K) map[K]int` | Counts elements per group determined by `fn`. |

### Random

| Function | Signature | Description |
|----------|-----------|-------------|
| **Shuffle** | `Shuffle[T any](x []T) []T` | Returns a randomly shuffled copy of the slice. |
| **Sample** | `Sample[T any](x []T, n int) []T` | Returns `n` random unique elements from the slice. |

### Utility

| Function | Signature | Description |
|----------|-----------|-------------|
| **Range** | `Range(args ...int) []int` | Generates a slice of integers. Accepts `(stop)`, `(start, stop)`, or `(start, stop, step)`. |

## Usage

```go
package main

import (
    "fmt"
    u "github.com/greggy/underscore-go"
)

func main() {
    nums := []int{1, 2, 3, 4, 5, 6}

    evens := u.Filter(nums, func(n int) bool { return n%2 == 0 })
    fmt.Println(evens) // [2 4 6]

    doubled := u.Map(nums, func(n int) int { return n * 2 })
    fmt.Println(doubled) // [2 4 6 8 10 12]

    sum := u.Reduce(nums, func(a, b int) int { return a + b }, 0)
    fmt.Println(sum) // 21

    grouped := u.GroupBy(nums, func(n int) string {
        if n%2 == 0 { return "even" }
        return "odd"
    })
    fmt.Println(grouped) // map[even:[2 4 6] odd:[1 3 5]]
}
```

## Development

```bash
make build   # compile
make test    # run tests
make lint    # go vet + gofmt check
make fmt-fix # auto-format
make check   # build + lint + test
```

## License

MIT
