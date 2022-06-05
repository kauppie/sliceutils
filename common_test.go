package sliceutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeSet(t *testing.T) {
	t.Run("make set from slice with repeating elements", func(t *testing.T) {
		slice := []int{1, 2, 3, 2, 4}
		set := makeSet(slice)

		assert.Equal(t, map[int]struct{}{
			1: {},
			2: {},
			3: {},
			4: {},
		}, set)
	})

	t.Run("Return empty map on empty slice", func(t *testing.T) {
		slice := []int{}
		set := makeSet(slice)

		assert.Equal(t, map[int]struct{}{}, set)
	})

	t.Run("Return empty map on nil slice", func(t *testing.T) {
		var slice []int = nil
		set := makeSet(slice)

		assert.Equal(t, map[int]struct{}{}, set)
	})
}

func TestZeroValue(t *testing.T) {
	t.Run("Return zero value for int", func(t *testing.T) {
		zero := zeroValue[int]()
		assert.Equal(t, 0, zero)
	})

	t.Run("Return zero value for string", func(t *testing.T) {
		zero := zeroValue[string]()
		assert.Equal(t, "", zero)
	})

	t.Run("Return zero value for pointer", func(t *testing.T) {
		zeroVal := zeroValue[*string]()
		var zeroPointer *string
		assert.Equal(t, zeroPointer, zeroVal)
	})

	t.Run("Return zero value for empty interface", func(t *testing.T) {
		zeroVal := zeroValue[any]()
		var emptyInterface interface{}
		assert.Equal(t, emptyInterface, zeroVal)
	})
}
