package sliceutils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

////////////////////////////////
//********** TESTS ***********//
////////////////////////////////

func TestAll(t *testing.T) {
	t.Run("All elements evaluate to true", func(t *testing.T) {
		slice := []int{1, 4, 6, 2, 3, 7}
		allPositive := All(slice, func(i int) bool { return i > 0 })
		assert.True(t, allPositive)
	})

	t.Run("Some elements don't evaluate to true", func(t *testing.T) {
		slice := []int{1, 4, 6, -2, 3, 7}
		allPositive := All(slice, func(i int) bool { return i > 0 })
		assert.False(t, allPositive)
	})

	t.Run("Return true on nil slice", func(t *testing.T) {
		var slice []int = nil
		allPositive := All(slice, func(i int) bool { return i > 0 })
		assert.True(t, allPositive)
	})
}

func TestAny(t *testing.T) {
	t.Run("Some elements evaluate to true", func(t *testing.T) {
		slice := []int{-1, -4, 6, -2, 3, 7}
		anyPositive := Any(slice, func(i int) bool { return i > 0 })
		assert.True(t, anyPositive)
	})

	t.Run("All elements evaluate to false", func(t *testing.T) {
		slice := []int{-1, -4, -6, -2, -3, -7}
		anyPositive := Any(slice, func(i int) bool { return i > 0 })
		assert.False(t, anyPositive)
	})

	t.Run("Return false on nil slice", func(t *testing.T) {
		var slice []int = nil
		anyPositive := Any(slice, func(i int) bool { return i > 0 })
		assert.False(t, anyPositive)
	})
}

func TestAreDisjoint(t *testing.T) {
	t.Run("Sets are disjoint", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{5, 4, 6}
		disjoint := AreDisjoint(a, b)
		assert.True(t, disjoint)
	})

	t.Run("Sets are not disjoint", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{3, 2, 6}
		disjoint := AreDisjoint(a, b)
		assert.False(t, disjoint)
	})

	t.Run("Nil sets are disjoint", func(t *testing.T) {
		disjoint := AreDisjoint[int](nil, nil)
		assert.True(t, disjoint)
	})
}

func TestContains(t *testing.T) {
	t.Run("Slice contains element", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}
		contains := Contains(slice, 2)
		assert.True(t, contains)
	})

	t.Run("Slice does not contain element", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}
		contains := Contains(slice, 5)
		assert.False(t, contains)
	})

	t.Run("Return false on nil slice", func(t *testing.T) {
		var slice []int = nil
		contains := Contains(slice, 1)
		assert.False(t, contains)
	})
}

func TestCount(t *testing.T) {
	t.Run("Count zeros", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 0, 1, 4, 0, 0, 12, 3, 5, 7, 1}
		count := Count(slice, func(i int) bool { return i == 0 })
		assert.Equal(t, 3, count)
	})

	t.Run("Return zero on nil slice", func(t *testing.T) {
		var slice []int = nil
		count := Count(slice, func(i int) bool { return i == 0 })
		assert.Equal(t, 0, count)
	})
}

func TestDeduplicate(t *testing.T) {
	t.Run("Slice with duplicates", func(t *testing.T) {
		slice := []int{1, 2, 3, 2}
		depupped := Deduplicate(slice)
		assert.Equal(t, []int{1, 2, 3}, depupped)
	})

	t.Run("Slice without duplicates", func(t *testing.T) {
		slice := []int{1, 2, 3}
		depupped := Deduplicate(slice)
		assert.Equal(t, []int{1, 2, 3}, depupped)
	})

	t.Run("Return nil on nil slice", func(t *testing.T) {
		var slice []int = nil
		depupped := Deduplicate(slice)
		assert.Nil(t, depupped)
	})
}

func TestDeduplicateInPlace(t *testing.T) {
	t.Run("Slice with duplicates", func(t *testing.T) {
		slice := []int{1, 2, 3, 2}
		DeduplicateInPlace(&slice)
		assert.Equal(t, []int{1, 2, 3}, slice)
	})

	t.Run("Slice without duplicates", func(t *testing.T) {
		slice := []int{1, 2, 3}
		DeduplicateInPlace(&slice)
		assert.Equal(t, []int{1, 2, 3}, slice)
	})

	t.Run("Return nil on nil slice", func(t *testing.T) {
		var slice []int = nil
		DeduplicateInPlace(&slice)
		assert.Nil(t, slice)
	})

	t.Run("Do nothing on nil slice pointer", func(t *testing.T) {
		DeduplicateInPlace[int](nil)
	})
}

func TestDifference(t *testing.T) {
	t.Run("Difference of two overlapping sets", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{3, 2, 6}
		difference := Difference(a, b)
		assert.Equal(t, []int{1}, difference)
	})

	t.Run("Difference of two non-overlapping sets", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{5, 4, 6}
		difference := Difference(a, b)
		assert.Equal(t, []int{1, 2, 3}, difference)
	})

	t.Run("Return nil when both sets are nil", func(t *testing.T) {
		difference := Difference[int](nil, nil)
		assert.Nil(t, difference)
	})
}

func TestFilter(t *testing.T) {
	t.Run("Retain strings shorter than 4 characters", func(t *testing.T) {
		slice := []string{"hello", "foo", "bar", "pointer", "cow", "F"}
		filtered := Filter(slice, func(s string) bool { return len(s) < 4 })
		assert.Equal(t, []string{"foo", "bar", "cow", "F"}, filtered)
	})

	t.Run("Return nil on nil slice", func(t *testing.T) {
		var slice []int = nil
		flat := Filter(slice, func(i int) bool { return i < 4 })
		assert.Nil(t, flat)
	})
}

func TestFilterInPlace(t *testing.T) {
	t.Run("Retain strings shorter than 4 characters", func(t *testing.T) {
		slice := []string{"hello", "foo", "bar", "pointer", "cow", "F"}
		FilterInPlace(&slice, func(s string) bool { return len(s) < 4 })
		assert.Equal(t, []string{"foo", "bar", "cow", "F"}, slice)
	})

	t.Run("Do nothing on nil slice", func(t *testing.T) {
		var slice []int = nil
		FilterInPlace(&slice, func(i int) bool { return i < 4 })
		assert.Nil(t, slice)
	})

	t.Run("Do nothing on nil slice pointer", func(t *testing.T) {
		FilterInPlace(nil, func(i int) bool { return i < 4 })
	})
}

func TestFilterMap(t *testing.T) {
	ToPointer := func(s string) *string {
		return &s
	}

	t.Run("Pointers to concrete types", func(t *testing.T) {
		slice := []*string{
			ToPointer("hello"),
			nil,
			ToPointer("foo"),
			nil,
			nil,
			ToPointer("bar"),
		}
		filterMapped := FilterMap(slice, func(s *string) (string, bool) {
			if s == nil {
				return "", false
			}
			return *s, true
		})
		assert.Equal(t, []string{"hello", "foo", "bar"}, filterMapped)
	})

	t.Run("Strings to their byte lengths if under 4", func(t *testing.T) {
		slice := []string{"hello", "foo", "bar", "pointer", "cow", "F"}
		filterMapped := FilterMap(slice, func(s string) (int, bool) {
			slen := len(s)
			return slen, slen < 4
		})
		assert.Equal(t, []int{3, 3, 3, 1}, filterMapped)
	})

	t.Run("Return nil on nil slice", func(t *testing.T) {
		var slice []int = nil
		filterMapped := FilterMap(slice, func(i int) (int, bool) { return i, i > 0 })
		assert.Nil(t, filterMapped)
	})
}

func TestFindBy(t *testing.T) {
	t.Run("Try to find and is found", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
		idx, found := FindBy(slice, func(i int) bool {
			return i == 6
		})
		assert.Equal(t, 5, idx)
		assert.True(t, found)
	})

	t.Run("Try to find and is not found", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
		idx, found := FindBy(slice, func(i int) bool {
			return i == 9
		})
		assert.Equal(t, 0, idx)
		assert.False(t, found)
	})

	t.Run("Return zero and false on nil slice", func(t *testing.T) {
		var slice []int = nil
		idx, found := FindBy(slice, func(i int) bool {
			return i == 0
		})
		assert.Equal(t, 0, idx)
		assert.False(t, found)
	})
}

func TestFlatten(t *testing.T) {
	t.Run("Flatten integer slice", func(t *testing.T) {
		slice := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}}
		flat := Flatten(slice)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, flat)
	})

	t.Run("Return nil on nil slice", func(t *testing.T) {
		var slice [][]int = nil
		flat := Flatten(slice)
		assert.Nil(t, flat)
	})
}

func TestFold(t *testing.T) {
	t.Run("Calculate sum and factorial", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		sum := Fold(numbers, 0, func(acc, next int) int {
			return acc + next
		})
		assert.Equal(t, 0+1+2+3+4+5+6, sum)

		factorial := Fold(numbers, 1, func(acc, next int) int {
			return acc * next
		})
		assert.Equal(t, 1*1*2*3*4*5*6, factorial)
	})

	t.Run("Return initial value on nil slice", func(t *testing.T) {
		var slice []int = nil

		folded := Fold(slice, 0, func(acc, next int) int {
			return acc + next
		})
		assert.Equal(t, 0, folded)

		folded2 := Fold(slice, 42, func(acc, next int) int {
			return acc + next
		})
		assert.Equal(t, 42, folded2)
	})
}

func TestFrequencies(t *testing.T) {
	t.Run("Count integer frequencies", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 0, 1, 4, 0, 0, 12, 3, 5, 7, 1}
		frequencies := Frequencies(slice)
		assert.Equal(t, map[int]int{
			1:  3,
			2:  1,
			3:  2,
			4:  2,
			0:  3,
			12: 1,
			5:  1,
			7:  1},
			frequencies,
		)
	})

	t.Run("Empty map on empty slice", func(t *testing.T) {
		slice := []int{}
		frequencies := Frequencies(slice)
		assert.Equal(t, map[int]int{}, frequencies)
	})

	t.Run("Return nil on nil slice", func(t *testing.T) {
		var slice []int = nil
		frequencies := Frequencies(slice)
		assert.Nil(t, frequencies)
	})
}

func TestIntersection(t *testing.T) {
	t.Run("Intersection of two overlapping sets", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{3, 2, 6}
		intersection := Intersection(a, b)
		assert.Equal(t, []int{2, 3}, intersection)
	})

	t.Run("Intersection of two non-overlapping sets", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{5, 4, 6}
		intersection := Intersection(a, b)
		assert.Equal(t, []int{}, intersection)
	})

	t.Run("Return nil when both sets are nil", func(t *testing.T) {
		intersection := Intersection[int](nil, nil)
		assert.Nil(t, intersection)
	})
}

func TestIsSet(t *testing.T) {
	t.Run("Is slice with only unique elements a set", func(t *testing.T) {
		set := []string{"foo", "bar", "hello", "world", "baz"}
		isSet := IsSet(set)
		assert.True(t, isSet)
	})

	t.Run("Is slice with repeating elements a set", func(t *testing.T) {
		notSet := []string{"foo", "bar", "baz", "foo", "hello"}
		isSet := IsSet(notSet)
		assert.False(t, isSet)
	})

	t.Run("Return true on nil slice", func(t *testing.T) {
		var set []string = nil
		isSet := IsSet(set)
		assert.True(t, isSet)
	})
}

func TestIsSortedBy(t *testing.T) {
	t.Run("Is sorted by with sorted slices", func(t *testing.T) {
		sortedSlice := []int{1, 2, 3, 4, 4, 5, 6, 7, 8}
		sorted := IsSortedBy(sortedSlice, func(lhs, rhs int) bool {
			return lhs < rhs
		})
		assert.True(t, sorted)

		sortedSlice2 := []string{"bar", "baz", "foo", "hello", "world"}
		sorted2 := IsSortedBy(sortedSlice2, func(lhs, rhs string) bool { return lhs < rhs })
		assert.True(t, sorted2)

		reverseSortedSlice := Reverse(sortedSlice)
		reverseSorted := IsSortedBy(reverseSortedSlice, func(lhs, rhs int) bool { return lhs > rhs })
		assert.True(t, reverseSorted)
	})

	t.Run("Is sorted by with unsorted slices", func(t *testing.T) {
		unsortedSlice := []int{1, 2, 3, 4, 5, 4, 6, 7, 8}
		sorted := IsSortedBy(unsortedSlice, func(lhs, rhs int) bool {
			return lhs < rhs
		})
		assert.False(t, sorted)

		unsortedSlice2 := []string{"baz", "bar", "foo", "hello", "world"}
		sorted2 := IsSortedBy(unsortedSlice2, func(lhs, rhs string) bool { return lhs < rhs })
		assert.False(t, sorted2)
	})
}

func TestIsSubSet(t *testing.T) {
	t.Run("IsSubSet on subset", func(t *testing.T) {
		super := []int{1, 2, 3}
		sub := []int{1, 2}
		assert.True(t, IsSubSet(sub, super))
	})

	t.Run("IsSubSet on non-subset", func(t *testing.T) {
		super := []int{1, 2, 3}
		notSub := []int{1, 2, 4}
		assert.False(t, IsSubSet(notSub, super))
	})

	t.Run("Equal set is subset of itself", func(t *testing.T) {
		set := []int{1, 2, 3}
		assert.True(t, IsSubSet(set, set))
	})

	t.Run("True on nil sets", func(t *testing.T) {
		assert.True(t, IsSubSet[int](nil, nil))
	})
}

func TestIsSuperSet(t *testing.T) {
	t.Run("IsSuperSet on superset", func(t *testing.T) {
		super := []int{1, 2, 3}
		sub := []int{1, 2}
		assert.True(t, IsSuperSet(super, sub))
	})

	t.Run("IsSuperSet on non-superset", func(t *testing.T) {
		notSuper := []int{1, 2, 3}
		sub := []int{1, 2, 4}
		assert.False(t, IsSuperSet(notSuper, sub))
	})

	t.Run("Equal set is superset of itself", func(t *testing.T) {
		set := []int{1, 2, 3}
		assert.True(t, IsSuperSet(set, set))
	})

	t.Run("True on nil sets", func(t *testing.T) {
		assert.True(t, IsSubSet[int](nil, nil))
	})
}

func TestJoin(t *testing.T) {
	t.Run("Join variadics", func(t *testing.T) {
		slice1 := []int{1, 2, 3}
		slice2 := []int{4, 5, 6}
		slice3 := []int{7, 8}

		joined := Join(slice1, slice2, slice3)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, joined)
	})

	t.Run("Join two dimensional slice", func(t *testing.T) {
		slice := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}}
		joined := Join(slice...)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, joined)
	})

	t.Run("Return empty slice on nil slices", func(t *testing.T) {
		var slice []int = nil
		var slice2 []int = nil
		joined := Join(slice, slice2)
		assert.Equal(t, []int{}, joined)
	})

	t.Run("Return nil on no arguments", func(t *testing.T) {
		joined := Join[int]()
		assert.Nil(t, joined)
	})
}

func TestMap(t *testing.T) {
	t.Run("Strings to their rune lengths", func(t *testing.T) {
		slice := []string{"bar", "", "f", "hello", "world"}
		lengths := Map(slice, func(s string) int { return len(s) })
		assert.Equal(t, []int{3, 0, 1, 5, 5}, lengths)
	})

	t.Run("Return nil on nil slice", func(t *testing.T) {
		var slice []string = nil
		outSlice := Map(slice, func(s string) int { return len(s) })
		assert.Nil(t, outSlice)
	})
}

func TestMaxBy(t *testing.T) {
	t.Run("Return max from slice", func(t *testing.T) {
		slice := []int{4, 5, 7, 3, 9, -1, 3, 4, 7, 12, 43, 10, 5}
		max, ok := MaxBy(slice, func(lhs, rhs int) bool {
			return lhs < rhs
		})
		assert.True(t, ok)
		assert.Equal(t, 43, max)
	})

	t.Run("Return zero value and false on empty slice", func(t *testing.T) {
		slice := []int{}
		max, ok := MaxBy(slice, func(lhs, rhs int) bool {
			return lhs < rhs
		})
		assert.False(t, ok)
		assert.Zero(t, max)
	})
}

func TestMinBy(t *testing.T) {
	t.Run("Return min from slice", func(t *testing.T) {
		slice := []int{4, 5, 7, 3, 9, -1, 3, 4, 7, 12, 43, 10, 5}
		min, ok := MinBy(slice, func(lhs, rhs int) bool {
			return lhs < rhs
		})
		assert.True(t, ok)
		assert.Equal(t, -1, min)
	})

	t.Run("Return zero value and false on empty slice", func(t *testing.T) {
		slice := []int{}
		min, ok := MinBy(slice, func(lhs, rhs int) bool {
			return lhs < rhs
		})
		assert.False(t, ok)
		assert.Zero(t, min)
	})
}

func TestPartition(t *testing.T) {
	t.Run("Partition by integer parity", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		even, odd := Partition(slice, func(i int) bool { return i%2 == 0 })
		assert.Equal(t, []int{2, 4, 6, 8, 10}, even)
		assert.Equal(t, []int{1, 3, 5, 7, 9}, odd)
	})

	t.Run("Return nil on nil slice", func(t *testing.T) {
		var slice []int = nil
		even, odd := Partition(slice, func(i int) bool { return i%2 == 0 })
		assert.Nil(t, even)
		assert.Nil(t, odd)
	})
}

func TestPartitionInPlace(t *testing.T) {
	t.Run("Partition with even number of elements", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		idx := PartitionInPlace(slice, func(i int) bool { return i%2 == 0 })
		assert.Equal(t, []int{10, 2, 8, 4, 6}, slice[:idx])
		assert.Equal(t, []int{5, 7, 3, 9, 1}, slice[idx:])
	})

	t.Run("Partition with odd number of elements", func(t *testing.T) {
		slice := []int{1, 3, 4, -1, -5, 10, 9, -4, -3}
		idx := PartitionInPlace(slice, func(i int) bool { return i > 0 })
		assert.Equal(t, []int{1, 3, 4, 9, 10}, slice[:idx])
		assert.Equal(t, []int{-5, -1, -4, -3}, slice[idx:])
	})

	t.Run("Partition with single element", func(t *testing.T) {
		slice := []int{1}
		idx := PartitionInPlace(slice, func(i int) bool { return i > 0 })
		assert.Equal(t, []int{1}, slice[:idx])
		assert.Equal(t, []int{}, slice[idx:])
	})

	t.Run("Do nothing on empty slice and return zero index", func(t *testing.T) {
		slice := []int{}
		idx := PartitionInPlace(slice, func(i int) bool { return i%2 == 0 })
		assert.Equal(t, 0, idx)
	})

	t.Run("Do nothing on nil slice and return zero index", func(t *testing.T) {
		var slice []int = nil
		idx := PartitionInPlace(slice, func(i int) bool { return i%2 == 0 })
		assert.Equal(t, 0, idx)
	})
}

func TestReverse(t *testing.T) {
	t.Run("Reverse integer slice", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		reversed := Reverse(slice)
		assert.Equal(t, []int{5, 4, 3, 2, 1}, reversed)
	})

	t.Run("Return nil on nil slice", func(t *testing.T) {
		var slice []int = nil
		output := Reverse(slice)
		assert.Nil(t, output)
	})
}

func TestReverseInPlace(t *testing.T) {
	t.Run("Reverse integer slice", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		ReverseInPlace(slice)
		assert.Equal(t, []int{5, 4, 3, 2, 1}, slice)
	})

	t.Run("Do nothing on nil slice", func(t *testing.T) {
		var slice []int = nil
		ReverseInPlace(slice)
		assert.Nil(t, slice)
	})
}

func TestSymmetricDifference(t *testing.T) {
	t.Run("Symmetric difference on two overlapping sets", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{3, 2, 6}
		symmDiff := SymmetricDifference(a, b)
		assert.Equal(t, []int{1, 6}, symmDiff)
	})

	t.Run("Preserve left set on empty right set", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{}
		symmDiff := SymmetricDifference(a, b)
		assert.Equal(t, []int{1, 2, 3}, symmDiff)
	})

	t.Run("Empty set on empty sets", func(t *testing.T) {
		a := []int{}
		b := []int{}
		symmDiff := SymmetricDifference(a, b)
		assert.Equal(t, []int{}, symmDiff)
	})

	t.Run("Return nil when both sets are nil", func(t *testing.T) {
		symmDiff := SymmetricDifference[int](nil, nil)
		assert.Nil(t, symmDiff)
	})
}

func TestUnion(t *testing.T) {
	t.Run("Union on two overlapping sets", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{3, 2, 6}
		union := Union(a, b)
		assert.Equal(t, []int{1, 2, 3, 6}, union)
	})

	t.Run("Preserve left set on empty right set", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{}
		union := Union(a, b)
		assert.Equal(t, []int{1, 2, 3}, union)
	})

	t.Run("Empty set on empty sets", func(t *testing.T) {
		a := []int{}
		b := []int{}
		union := Union(a, b)
		assert.Equal(t, []int{}, union)
	})

	t.Run("Return nil when both sets are nil", func(t *testing.T) {
		union := Union[int](nil, nil)
		assert.Nil(t, union)
	})
}

////////////////////////////////
//******** BENCHMARKS ********//
////////////////////////////////

func BenchmarkAll(b *testing.B) {
	slice := []string{"boo", "bar", "baz", "hib", "heb", "obe", "lob", "suber",
		"library", "functional function", "slice", "NOW", "hey"}

	b.Run("Do all strings contain rune", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var _ = All(slice, func(x string) bool { return strings.ContainsRune(x, rune('b')) })
		}
	})
}

func BenchmarkAny(b *testing.B) {
	slice := []string{"foo", "bar", "baz", "his", "her", "one", "log", "super",
		"library", "functional function", "slice", "NOW", "hey"}

	b.Run("Does any string contain rune", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var _ = Any(slice, func(x string) bool { return strings.ContainsRune(x, rune('W')) })
		}
	})
}
