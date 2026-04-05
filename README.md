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

## Examples

### Each

```go
u.Each([]string{"a", "b", "c"}, func(s string) {
    fmt.Println(s)
})
// a
// b
// c
```

### Map

```go
u.Map([]int{1, 2, 3}, func(n int) int { return n * 2 })
// [2, 4, 6]

u.Map([]int{1, 2, 3}, func(n int) string { return fmt.Sprintf("#%d", n) })
// ["#1", "#2", "#3"]
```

### Filter

```go
u.Filter([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 })
// [2, 4]
```

### Reject

```go
u.Reject([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 })
// [1, 3, 5]
```

### Reduce

```go
u.Reduce([]int{1, 2, 3, 4}, func(a, b int) int { return a + b }, 0)
// 10
```

### ReduceRight

```go
u.ReduceRight([]string{"a", "b", "c"}, func(a, b string) string { return a + b }, "")
// "cba"
```

### Every

```go
u.Every([]int{2, 4, 6}, func(n int) bool { return n%2 == 0 })
// true

u.Every([]int{2, 3, 6}, func(n int) bool { return n%2 == 0 })
// false
```

### Some

```go
u.Some([]int{1, 3, 5, 6}, func(n int) bool { return n%2 == 0 })
// true

u.Some([]int{1, 3, 5}, func(n int) bool { return n%2 == 0 })
// false
```

### Partition

```go
pass, fail := u.Partition([]int{1, 2, 3, 4, 5}, func(n int) bool { return n > 3 })
// pass: [4, 5]
// fail: [1, 2, 3]
```

### Contains

```go
u.Contains([]string{"a", "b", "c"}, "b")
// true

u.Contains([]string{"a", "b", "c"}, "z")
// false
```

### Find

```go
result := u.Find([]int{1, 2, 3, 4}, func(n int) bool { return n > 2 })
// *result == 3

result = u.Find([]int{1, 2, 3}, func(n int) bool { return n > 10 })
// result == nil
```

### FindIndex

```go
u.FindIndex([]int{10, 20, 30}, func(n int) bool { return n >= 20 })
// 1

u.FindIndex([]int{10, 20, 30}, func(n int) bool { return n > 100 })
// -1
```

### FindLastIndex

```go
u.FindLastIndex([]int{1, 2, 3, 2, 1}, func(n int) bool { return n == 2 })
// 3
```

### IndexOf

```go
u.IndexOf([]string{"a", "b", "c", "b"}, "b")
// 1

u.IndexOf([]string{"a", "b", "c"}, "z")
// -1
```

### LastIndexOf

```go
u.LastIndexOf([]int{1, 2, 3, 2, 1}, 2)
// 3
```

### First

```go
u.First([]int{1, 2, 3, 4, 5}, 3)
// [1, 2, 3]
```

### Last

```go
u.Last([]int{1, 2, 3, 4, 5}, 2)
// [4, 5]
```

### Initial

```go
u.Initial([]int{1, 2, 3, 4, 5}, 2)
// [1, 2, 3]
```

### Rest

```go
u.Rest([]int{1, 2, 3, 4, 5}, 2)
// [3, 4, 5]
```

### Flatten

```go
u.Flatten([][]int{{1, 2}, {3, 4}, {5}})
// [1, 2, 3, 4, 5]
```

### Chunk

```go
u.Chunk([]int{1, 2, 3, 4, 5}, 2)
// [[1, 2], [3, 4], [5]]
```

### Uniq

```go
u.Uniq([]int{1, 2, 1, 3, 2, 4})
// [1, 2, 3, 4]
```

### Union

```go
u.Union([]int{1, 2, 3}, []int{3, 4, 5})
// [1, 2, 3, 4, 5]
```

### Intersection

```go
u.Intersection([]int{1, 2, 3, 4}, []int{2, 4, 6}, []int{4, 2, 8})
// [2, 4]
```

### Difference

```go
u.Difference([]int{1, 2, 3, 4, 5}, []int{2, 4}, []int{5})
// [1, 3]
```

### Without

```go
u.Without([]int{1, 2, 3, 4, 5}, 3, 5)
// [1, 2, 4]
```

### Min

```go
u.Min([]int{3, 1, 4, 1, 5})
// 1

u.Min([]float64{2.7, 1.1, 3.5})
// 1.1
```

### Max

```go
u.Max([]int{3, 1, 4, 1, 5})
// 5
```

### SortBy

```go
u.SortBy([]int{3, 1, 2}, func(n int) int { return n })
// [1, 2, 3]

u.SortBy([]int{-3, 1, -2}, func(n int) int { return n * n })
// [1, -2, -3]
```

### GroupBy

```go
u.GroupBy([]string{"one", "two", "three"}, func(s string) int { return len(s) })
// map[3:["one" "two"] 5:["three"]]
```

### CountBy

```go
u.CountBy([]int{1, 2, 3, 4, 5}, func(n int) string {
    if n%2 == 0 { return "even" }
    return "odd"
})
// map["even":2 "odd":3]
```

### Shuffle

```go
u.Shuffle([]int{1, 2, 3, 4, 5})
// [3, 5, 1, 4, 2] (random order)
```

### Sample

```go
u.Sample([]int{1, 2, 3, 4, 5}, 3)
// [5, 2, 4] (3 random unique elements)
```

### Range

```go
u.Range(5)
// [0, 1, 2, 3, 4]

u.Range(1, 5)
// [1, 2, 3, 4]

u.Range(0, 20, 5)
// [0, 5, 10, 15]

u.Range(0, -5, -1)
// [0, -1, -2, -3, -4]
```

## Development

```bash
make build   # compile
make test    # run tests
make lint    # golangci-lint
make fmt-fix # auto-format
make check   # build + lint + test
```

## License

MIT
