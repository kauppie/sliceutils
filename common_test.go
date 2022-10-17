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

func TestSliceDivGen(t *testing.T) {
	type expectedOut struct {
		offset int
		length int
	}

	type testCase struct {
		name    string
		gen     sliceDivGen
		expects []expectedOut
	}

	testCases := []testCase{
		{
			name: "Length 10 and divs 2",
			gen:  newSliceDivGen(10, 2),
			expects: []expectedOut{
				{
					offset: 0,
					length: 5,
				},
				{
					offset: 5,
					length: 5,
				},
			},
		},
		{
			name: "Length 10 and divs 3",
			gen:  newSliceDivGen(10, 3),
			expects: []expectedOut{
				{
					offset: 0,
					length: 4,
				},
				{
					offset: 4,
					length: 3,
				},
				{
					offset: 7,
					length: 3,
				},
			},
		},
		{
			name: "Length 10 and divs 4",
			gen:  newSliceDivGen(10, 4),
			expects: []expectedOut{
				{
					offset: 0,
					length: 3,
				},
				{
					offset: 3,
					length: 3,
				},
				{
					offset: 6,
					length: 2,
				},
				{
					offset: 8,
					length: 2,
				},
			},
		},
		{
			name: "Length 4 and divs 4",
			gen:  newSliceDivGen(4, 4),
			expects: []expectedOut{
				{
					offset: 0,
					length: 1,
				},
				{
					offset: 1,
					length: 1,
				},
				{
					offset: 2,
					length: 1,
				},
				{
					offset: 3,
					length: 1,
				},
			},
		},
		{
			name: "Length 1 and divs 4",
			gen:  newSliceDivGen(1, 4),
			expects: []expectedOut{
				{
					offset: 0,
					length: 1,
				},
				{
					offset: 1,
					length: 0,
				},
				{
					offset: 1,
					length: 0,
				},
				{
					offset: 1,
					length: 0,
				},
			},
		},
		{
			name: "Length 0 and divs 4",
			gen:  newSliceDivGen(0, 4),
			expects: []expectedOut{
				{
					offset: 0,
					length: 0,
				},
				{
					offset: 0,
					length: 0,
				},
				{
					offset: 0,
					length: 0,
				},
				{
					offset: 0,
					length: 0,
				},
			},
		},
		{
			name: "Length 5 and divs 1",
			gen:  newSliceDivGen(5, 1),
			expects: []expectedOut{
				{
					offset: 0,
					length: 5,
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			for i, expected := range testCase.expects {
				offset, length := testCase.gen.offsetAndLength(i)
				assert.Equal(t, expected.offset, offset)
				assert.Equal(t, expected.length, length)
			}
		})
	}
}
