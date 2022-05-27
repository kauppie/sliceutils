package sliceutils

// Creates a set out of slice elements. Duplicates are discarded.
func makeSet[T comparable](slice []T) map[T]struct{} {
	uniques := make(map[T]struct{})
	for _, val := range slice {
		uniques[val] = struct{}{}
	}
	return uniques
}
