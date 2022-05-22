# **_SliceUtils | Slice Utility Functions_**

[![CI status](https://github.com/kauppie/sliceutils/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/kauppie/sliceutils/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kauppie/sliceutils)](https://goreportcard.com/report/github.com/kauppie/sliceutils)
[![codecov](https://codecov.io/gh/kauppie/sliceutils/branch/main/graph/badge.svg)](https://codecov.io/gh/kauppie/sliceutils)
[![MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/kauppie/sliceutils/blob/main/LICENSE)

This library implements several high-level functions useful for interacting with slices. It reduces boilerplate required by for-loops and variable initializations.

Library takes advantage of Go generics increasing usability by compile-time type-safety. Go version of at least **1.18** is therefore required.

## Functions

### >> _All_

Returns true if all slice elements are evaluated true with given argument function.

### >> _Any_

Returns true if any of slice element is evaluated true with given argument function.

### >> _Count_

Counts the number of elements in a slice for which the argument function returns true.

### >> _Filter_

Creates a slice where slice elements for which the argument function returns false are filtered out.

### >> _FilterMap_

Filters _and_ maps slice elements to new slice. See [_Filter_](#filter) and [_Map_](#map) for more details. This function exists to allow better performance than using _Filter_ and _Map_ separately.

### >> _FindBy_

Searches to find element's index in a slice for which the argument function returns true.

### >> _Flatten_

Converts a _N_-dimensional slice into a _N-1_ -dimensional slice.

### >> _Fold_

Folds a slice into a single value. Other name for such a function is _reduce_.

It starts with a initial value and updates it iteratively using the argument function and slice's elements to accumulate the final result.

### >> _Frequencies_

Counts the number of occurrences for each element. Requires slice elements to be `comparable`.

### >> _IsSortedBy_

Returns true for slices whose elements are sorted according to passed argument function.

### >> _IsSet_

Returns true for slices that are sets i.e. contain only unique elements. Requires slice elements to be `comparable`.

### >> _Join_

Joins one or more slices together. Similar to [_Flatten_](#flatten) but uses variadic arguments instead.

### >> _Map_

Maps each element through argument function which can modify their type and/or value.

### >> _MaxBy_

Returns the maximum element value in a slice using provided comparison function.

### >> _MinBy_

Returns the minimum element value in a slice using provided comparison function.

### >> _Partition_

Partitions slice elements into two separate slices by argument function's boolean return value.

### >> _Reverse_

Reverses the order of elements in a slice.

## Performance

Currently all the functions have at most **O(n \* m)** time complexity, where **n** is length of the argument slice and **m** is time complexity of the argument function. Functions without argument functions have time complexity of at most **O(n)**.

Performance over traditional for-loops is not _yet_ tested.
