# **_SliceUtils | Slice Utility Functions_**

[![CI status](https://github.com/kauppie/sliceutils/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/kauppie/sliceutils/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kauppie/sliceutils)](https://goreportcard.com/report/github.com/kauppie/sliceutils)
[![codecov](https://codecov.io/gh/kauppie/sliceutils/branch/main/graph/badge.svg)](https://codecov.io/gh/kauppie/sliceutils)
[![MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/kauppie/sliceutils/blob/main/LICENSE)

This library implements several high-level functions useful for interacting with slices. It reduces boilerplate required by for-loops and variable initializations.

Library takes advantage of Go generics increasing usability by compile-time type-safety. Go version of at least **1.18** is therefore required.

## Examples

Below are few examples on how this library can make your code **more** concise and **less** error-prone.

### _Map strings to their lengths_

```go
var strings []string

// Replace
lens := make([]int, 0)
for _, s := range strings {
  lens = append(lens, len(s))
}

// With
lens := Map(strings, func(s string) int { return len(s) })
```

### _Get positive numbers from a slice_

```go
var nums []int

// Replace
pos := make([]int, 0)
for _, n := range nums {
  if n > 0 {
    pos = append(pos, n)
  }
}

// With
pos := Filter(nums, func(n int) bool { return n > 0 })
```

### _Flatten slice_

```go
var slices [][]int

// Replace
flat := make([]int, 0)
for _, slice := range slices {
  flat = append(flat, slice...)
}

// With
flat := Flatten(slices)
```

### _Deduplicate slice_

```go
var slice []int

// Replace
uniques := make(map[int]struct{})
dedup := make([]int, 0)
for _, val := range slice {
  if _, ok := uniques[val]; !ok {
    dedup = append(dedup, val)
    uniques[val] = struct{}{}
  }
}

// With
dedup := Deduplicate(slice)
```

## List of all functions

### >> _All_

Returns `true` if all slice elements are evaluated `true` with given argument function.

### >> _Any_

Returns `true` if any slice element is evaluated `true` with given argument function.

### >> _AreDisjoint_

Returns `true` if two slice sets do not have common elements.

### >> _Contains_

Returns `true` if slice contains given element.

### >> _Count_

Counts the number of elements in a slice for which the argument function returns `true`.

### >> _Deduplicate_

Removes duplicate elements from a slice creating a new slice.

### >> _DeduplicateInPlace_

Removes duplicate elements from a slice in place.

### >> _Difference_

Calculates a difference set between two slice sets.

### >> _Filter_

Creates a slice which contains slice elements for which the argument function returns `true`.

### >> _FilterInPlace_

Retains elements in a slice for which the argument function returns `true`. Modifies the original slice and therefore does not allocate.

### >> _FilterMap_

Filters _and_ maps slice elements to new slice. See [_Filter_](#filter) and [_Map_](#map) for more details. This function exists to allow better performance than using _Filter_ and _Map_ separately.

### >> _FindBy_

Searches to find element's index in a slice for which the argument function returns `true`.

### >> _Flatten_

Converts a _N_-dimensional slice into a _N-1_ -dimensional slice.

### >> _Fold_

Folds a slice into a single value. Other name for such a function is _reduce_.

It starts with a initial value and updates it iteratively using the argument function and slice's elements to accumulate the final result.

### >> _Frequencies_

Counts the number of occurrences for each element. Requires slice elements to be `comparable`.

### >> _Intersection_

Calculates a intersection set between two slice sets.

### >> _IsSet_

Returns `true` for slices that are sets i.e. contain only unique elements. Requires slice elements to be `comparable`.

### >> _IsSortedBy_

Returns `true` for slices whose elements are sorted according to passed argument function.

### >> _IsSubSet_

Returns `true` if first slice set is a subset of the second slice set.

### >> _IsSuperSet_

Returns `true` if first slice set is a superset of the second slice set.

### >> _Join_

Joins one or more slices together. Similar to [_Flatten_](#flatten) but uses variadic arguments instead.

### >> _Map_

Maps each element through argument function which can modify their type and/or value.

### >> _MapInPlace_

Maps each slice element to a new value of the same type with provided mapping function. Does the operation in place modifying the original slice.

### >> _MaxBy_

Returns the maximum element value in a slice using provided comparison function.

### >> _MinBy_

Returns the minimum element value in a slice using provided comparison function.

### >> _Partition_

Partitions slice elements into two separate slices by argument function's boolean return value.

### >> _PartitionInPlace_

Partitions a slice in place so that the first partition contains elements for which the argument function return `true`, and the second partition contains elements that the function returns `false` for.

### >> _Reverse_

Creates a slice where the order of elements are reversed.

### >> _ReverseInPlace_

Reverses the order of elements in a slice.

### >> _SymmetricDifference_

Calculates a symmetric difference set from two slice sets.

### >> _Union_

Calculates a union set from two slice sets.

## Performance

Currently all the functions have at most **O(n \* m)** time complexity, where **n** is length of the argument slice and **m** is time complexity of the argument function. Functions without argument functions have time complexity of at most **O(n)**.

Performance over traditional for-loops is not _yet_ tested.
