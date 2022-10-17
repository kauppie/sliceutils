package sliceutils

// Creates a set out of slice elements. Duplicates are discarded.
func makeSet[T comparable](slice []T) map[T]struct{} {
	uniques := make(map[T]struct{})
	for _, val := range slice {
		uniques[val] = struct{}{}
	}
	return uniques
}

// Returns the zero value of type T.
func zeroValue[T any]() T {
	var t T
	return t
}

// Slice division generator is used to evenly divide a slice into sub-slices
// which could be processed in parallel. All sub-slices are non-overlapping.
type sliceDivGen struct {
	// Minimum number of elements per division.
	minDivLen int
	// Number of divisions which have `minDivLen + 1` elements.
	firstPartDivs int
}

// Creates a new slice division generator. Takes parameter `length` as length
// of the slice and `divs` as the number times to divide the slice.
//
// Panics if `divs` is zero.
func newSliceDivGen(length, divs int) sliceDivGen {
	return sliceDivGen{
		minDivLen:     length / divs,
		firstPartDivs: length % divs,
	}
}

// Gets the offset in the original slice and length of the sub-slice for given
// sub-slice index.
//
// `divIdx` is expected to be less than the number of divisions.
func (sdg sliceDivGen) offsetAndLength(divIdx int) (int, int) {
	if divIdx < sdg.firstPartDivs {
		offset := (sdg.minDivLen + 1) * divIdx
		return offset, sdg.minDivLen + 1
	} else {
		offset := divIdx*sdg.minDivLen + sdg.firstPartDivs
		return offset, sdg.minDivLen
	}
}

// Gets start and end indexes in the original slice for given sub-slice index.
//
// `divIdx` is expected to be less than the number of divisions.
func (sdg sliceDivGen) startAndEnd(divIdx int) (int, int) {
	offset, length := sdg.offsetAndLength(divIdx)
	return offset, offset + length
}
