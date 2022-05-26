package sliceutils

// Returns true if all slice elements are evaluated true with given evaluator
// function.
//
// Returns true on nil slice. Panics on nil evaluator function.
func All[T any](slice []T, allFn func(T) bool) bool {
	for _, val := range slice {
		if !allFn(val) {
			return false
		}
	}
	return true
}

// Returns true if any slice element is evaluated true with given evaluator
// function.
//
// Returns false on nil slice. Panics on nil evaluator function.
func Any[T any](slice []T, anyFn func(T) bool) bool {
	for _, val := range slice {
		if anyFn(val) {
			return true
		}
	}
	return false
}

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

// Filter values in a slice in place by filter function. Modified slice will
// contain values for which the filter function returns true. Slice is passed
// as pointer because its length could be modified.
//
// Does not allocate. Panics on nil filter function.
func FilterInPlace[T any](slicep *[]T, filterFn func(T) bool) {
	// Pointer could be nil.
	if slicep == nil {
		return
	}
	n := 0
	for _, val := range *slicep {
		if filterFn(val) {
			(*slicep)[n] = val
			n++
		}
	}
	// Possibly shorten the slice to current length.
	*slicep = (*slicep)[:n]
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

// Returns index of the found element and true in a tuple. If element is not
// found, returns zero and false.
//
// Returns zero and false on nil slice. Panics on nil find function.
func FindBy[T any](slice []T, findFn func(T) bool) (int, bool) {
	for i, val := range slice {
		if findFn(val) {
			return i, true
		}
	}
	return 0, false
}

// Flattens a N-dimensional slice to a N-1 -dimensional slice. Resulting slice
// preserves order from the original slice where the first values will be from
// the first slice.
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

// Folds a slice successively into single value. `init` is the initial value
// for which the fold function is applied. Fold function takes the current
// folded value and the next slice value and returns the folded value.
//
// Return initial value on nil slice. Panics on nil fold function.
func Fold[T, U any](slice []T, init U, foldFn func(U, T) U) U {
	for _, val := range slice {
		init = foldFn(init, val)
	}
	return init
}

// Returns the frequency of values in a slice. Resulting map contains the found
// values as keys and their number of occurrences as values.
//
// Returns nil on nil slice.
func Frequencies[T comparable](slice []T) map[T]int {
	// Preserve nil.
	if slice == nil {
		return nil
	}
	outMap := make(map[T]int)
	for _, val := range slice {
		// Missing value returns default which is zero.
		outMap[val] = outMap[val] + 1
	}
	return outMap
}

// Returns true if the slice is sorted by given comparison function. For
// ascending order, pass a comparison function which returns true when left is
// less than right.
//
// Returns true on nil slice. Panics on nil comparison function.
func IsSortedBy[T any](slice []T, lessFn func(T, T) bool) bool {
	for i := 1; i < len(slice); i++ {
		if lessFn(slice[i], slice[i-1]) {
			return false
		}
	}
	return true
}

// Returns true if the slice is a set e.g. contains only unique elements.
//
// Returns true on nil slice.
func IsSet[T comparable](slice []T) bool {
	uniques := make(map[T]struct{})
	for _, val := range slice {
		if _, found := uniques[val]; found {
			return false
		} else {
			uniques[val] = struct{}{}
		}
	}
	return true
}

// Join multiple slices together into a single slice. This is a variadic
// version of Flatten. The effective difference between Join and Flatten is
// that this returns empty slice on nil slice arguments while Flatten returns
// nil slice.
//
// Returns nil on no arguments. Returns empty slice on nil slice arguments.
func Join[T any](slices ...[]T) []T {
	// Preserve nil if no arguments.
	if slices == nil {
		return nil
	}
	outSlice := make([]T, 0)
	for _, slice := range slices {
		outSlice = append(outSlice, slice...)
	}
	return outSlice
}

// Maps each slice value with mapping function. Resulting slice contains values
// returned by the mapping function while preserving order.
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

// Returns the maximum element value and true from non-empty slice using
// the provided comparison function. To get maximum value, pass a comparison
// function which returns true when left is less than right. Function is
// stable, i.e. returns the first occurrence of maximum value.
//
// If slice is empty, returns zero value of type T and false.
func MaxBy[T any](slice []T, lessFn func(T, T) bool) (T, bool) {
	if len(slice) == 0 {
		var t T
		return t, false
	}
	max := slice[0]
	for _, val := range slice[1:] {
		if lessFn(max, val) {
			max = val
		}
	}
	return max, true
}

// Returns the minimum element value and true from non-empty slice using
// the provided comparison function. To get minimum value, pass a comparison
// function which returns true when left is less than right. Function is
// stable, i.e. returns the first occurrence of minimum value.
//
// If slice is empty, returns zero value of type T and false.
func MinBy[T any](slice []T, lessFn func(T, T) bool) (T, bool) {
	if len(slice) == 0 {
		var t T
		return t, false
	}
	min := slice[0]
	for _, val := range slice[1:] {
		if lessFn(val, min) {
			min = val
		}
	}
	return min, true
}

// Partition single slice into two slices using partition function. The first
// returned slice contains values for which the partition function returns true,
// and the second slice values for which the function returns false.
//
// Returns nil slices on nil slice. Panics on nil partition function.
func Partition[T any](slice []T, firstPart func(T) bool) ([]T, []T) {
	// Preserve nil.
	if slice == nil {
		return nil, nil
	}
	trueSlice := make([]T, 0)
	falseSlice := make([]T, 0)
	for _, val := range slice {
		if firstPart(val) {
			trueSlice = append(trueSlice, val)
		} else {
			falseSlice = append(falseSlice, val)
		}
	}
	return trueSlice, falseSlice
}

// Partition slice in place using partition function. First part contains
// elements for which the partition function returns true, and the second part
// values for which the function returns false. Function returns the index
// of the first element in the second partition.
//
// Does not allocate. Panics on nil partition function.
func PartitionInPlace[T any](slice []T, firstPart func(T) bool) int {
	fwd := -1
	bck := len(slice)
	for {
		// Increment indexes until elements at wrong partitions are found.
		fwd++
		for fwd < bck && firstPart(slice[fwd]) {
			fwd++
		}
		bck--
		for fwd < bck && !firstPart(slice[bck]) {
			bck--
		}
		if fwd >= bck {
			return fwd
		}
		// Swap elements in wrong partitions.
		slice[fwd], slice[bck] = slice[bck], slice[fwd]
	}
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

// Reverses the order of elements in a slice.
//
// Does not allocate.
func ReverseInPlace[T any](slice []T) {
	l := len(slice)
	for i := 0; i < l/2; i++ {
		slice[i], slice[l-1-i] = slice[l-1-i], slice[i]
	}
}
