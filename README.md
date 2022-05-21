# ***SliceUtils | Slice Utility Functions***

[![CI status](https://github.com/kauppie/sliceutils/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/kauppie/sliceutils/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kauppie/sliceutils)](https://goreportcard.com/report/github.com/kauppie/sliceutils)
[![codecov](https://codecov.io/gh/kauppie/sliceutils/branch/main/graph/badge.svg)](https://codecov.io/gh/kauppie/sliceutils)
[![MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/kauppie/sliceutils/blob/main/LICENSE)

This library implements several high-level functions useful for interacting with slices. It reduces boilerplate required by for-loops and variable initializations. 

Library takes advantage of Go generics increasing usability by compile-time type-safety. Go version of at least **1.18** is therefore required.

## Functions

### >> *Count*

Function counts the number of elements in a slice for which the argument function returns true.

### >> *Filter*

Function creates a slice where slice elements for which the argument function returns false are filtered out.

### >> *FilterMap*

Function filters *and* maps slice elements to new slice. See [*Filter*](#filter) and [*Map*](#map) for more details. This function exists to allow better performance than using *Filter* and *Map* separately.

### >> *Flatten*

Function converts a *N*-dimensional slice into a *N-1* -dimensional slice.

### >> *Fold*

Function folds a slice into a single value. Other name for such a function is *reduce*.

It starts with a initial value and updates it iteratively using the argument function and slice's elements to accumulate the final result.

### >> *Frequencies*

Function counts the number of occurrences for each unique element. Requires slice elements to be `comparable`.

### >> *IsSortedBy*

Function returns true for slices whose elements are sorted according to passed argument function.

### >> *IsSet*

Function returns true for slices that are sets. Requires slice elements to be `comparable`.

### >> *Map*

Function maps each element through argument function which can modify their type and/or value.

### >> *Partition*

Function partitions slice elements into two separate slices by argument function's boolean return value.

### >> *Reverse*

Function reverses the order of elements in a slice.

## Performance

Currently all the functions have at most **O(n * m)** time complexity, where **n** is length of the argument slice and **m** is time complexity of the argument function. Functions without argument functions have time complexity of at most **O(n)**.

Performance over traditional for-loops is not *yet* tested.
