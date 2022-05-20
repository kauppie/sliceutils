package sliceutils

// Count the number of matching items in a slice. Counter is incremented if
// counter function returns true on them.
//
// Panics on nil counter function.
func Count[T any](slice []T, counterFn func(T) bool) int {
	count := 0
	for _, val := range slice {
		if counterFn(val) {
			count++
		}
	}
	return count
}

// Filter values in a slice by filter function. Resulting slice will contain
// values for which the filter function returns true.
//
// Returns nil on nil slice. Panics on nil filter function.
func Filter[T any](slice []T, filterFn func(T) bool) []T {
	// Preserve nil.
	if slice == nil {
		return nil
	}
	outSlice := make([]T, 0)
	for _, val := range slice {
		if filterFn(val) {
			outSlice = append(outSlice, val)
		}
	}
	return outSlice
}

// Filter and map slice values with filter map function. Resulting slice
// will contain mapped values for which the filter map function returns true as
// the second argument. FilterMap is usually more efficient than using Filter
// and Map separately.
//
// Returns nil on nil slice. Panics on nil filter map function.
func FilterMap[T, U any](slice []T, filterMapFn func(T) (U, bool)) []U {
	// Preserve nil.
	if slice == nil {
		return nil
	}
	outSlice := make([]U, 0)
	for _, val := range slice {
		if mapped, ok := filterMapFn(val); ok {
			outSlice = append(outSlice, mapped)
		}
	}
	return outSlice
}

// Flatten a slice of slices into a slice. Resulting slice preserves order
// from the original slice where the first values will be from the first slice.
//
// Returns nil on nil slice.
func Flatten[T any](slice [][]T) []T {
	// Preserve nil.
	if slice == nil {
		return nil
	}
	outSlice := make([]T, 0)
	for _, val := range slice {
		outSlice = append(outSlice, val...)
	}
	return outSlice
}

// Returns the frequency of values in a slice. Resulting map contains the found
// values as keys and their number of occurrences as values.
//
// Returns nil on nil slice.
func Frequency[T comparable](slice []T) map[T]int {
	// Preserve nil.
	if slice == nil {
		return nil
	}
	outMap := make(map[T]int)
	for _, val := range slice {
		if _, exists := outMap[val]; exists {
			outMap[val]++
		} else {
			outMap[val] = 1
		}
	}
	return outMap
}

// Returns true if the slice is sorted by given comparison function. For
// ascending order, pass a comparison function which returns true when left is
// less than right.
//
// Returns true on nil slice. Panics on nil comparison function.
func IsSortedBy[T any](slice []T, compFn func(T, T) bool) bool {
	for i := 1; i < len(slice); i++ {
		if compFn(slice[i], slice[i-1]) {
			return false
		}
	}
	return true
}

// Maps each slice value with mapping function.
//
// Returns nil on nil slice. Panics on nil map function.
func Map[T, U any](slice []T, mapFn func(T) U) []U {
	// Preserve nil.
	if slice == nil {
		return nil
	}
	// Reserve capacity eagerly to allocate only once.
	outSlice := make([]U, 0, len(slice))
	for _, val := range slice {
		outSlice = append(outSlice, mapFn(val))
	}
	return outSlice
}

// Partition single slice into two slices using partition function. The first
// returned slice contains values for which the partition function returns true,
// and the second slice values for which the function returns false.
//
// Returns nil slices on nil slice. Panics on nil map function.
func Partition[T any](slice []T, partFn func(T) bool) ([]T, []T) {
	// Preserve nil.
	if slice == nil {
		return nil, nil
	}
	trueSlice := make([]T, 0)
	falseSlice := make([]T, 0)
	for _, val := range slice {
		if partFn(val) {
			trueSlice = append(trueSlice, val)
		} else {
			falseSlice = append(falseSlice, val)
		}
	}
	return trueSlice, falseSlice
}

// Reduces a slice successively into a single value. init is the initial value
// for which the reducer function is applied. Reducer function takes the current
// reduced value and the next slice value and returns the reduced value.
//
// Return initial value on nil slice. Panics on nil reducer function.
func Reduce[T, U any](slice []T, init U, reducerFn func(U, T) U) U {
	for _, val := range slice {
		init = reducerFn(init, val)
	}
	return init
}

// Reverses the order of elements in a slice.
//
// Returns nil on nil slice.
func Reverse[T any](slice []T) []T {
	// Preserve nil.
	if slice == nil {
		return nil
	}
	i := len(slice) - 1
	outSlice := make([]T, len(slice))
	for _, val := range slice {
		outSlice[i] = val
		i--
	}
	return outSlice
}
