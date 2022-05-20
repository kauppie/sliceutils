package sliceutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
