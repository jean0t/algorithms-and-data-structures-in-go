package tests

import (
    "testing"
    "reflect"
    "github.com/jean0t/algorithms-and-data-structures-in-go/algorithms/sorting"

)

type testCase struct {
    name string
    input []int
    expected []int
}

var testCases []testCase = []testCase{
    {
		name:     "Empty array",
		input:    []int{},
		expected: []int{},
	},
	{
		name:     "Single element array",
		input:    []int{5},
		expected: []int{5},
	},
	{
		name:     "Already sorted array",
		input:    []int{1, 2, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "Reverse sorted array",
		input:    []int{5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "Array with duplicates",
		input:    []int{3, 3, 2, 7, 7, 1, 2},
		expected: []int{1, 2, 2, 3, 3, 7, 7},
	},
	{
		name:     "Shuffled array",
		input:    []int{5, 2, 7, 1, 9, 3},
		expected: []int{1, 2, 3, 5, 7, 9},
	},
	{
		name:     "All elements same",
		input:    []int{4, 4, 4, 4},
		expected: []int{4, 4, 4, 4},
	},
	{
		name:     "Negative numbers",
		input:    []int{-3, 0, -1, 2},
		expected: []int{-3, -1, 0, 2},
	},
}


func TestQuickSort(t *testing.T) {
    for _, tt := range testCases {
        t.Run(tt.name, func(t *testing.T) {
            actual := make([]int, len(tt.input))
            copy(actual, tt.input)

            sorting.QuickSort(actual)

            if !reflect.DeepEqual(actual, tt.expected) {
                t.Errorf("Quicksort(%v) = %v; want %v\n", tt.input, actual, tt.expected)
            }
        })
    }
}



func TestRandomizedQuickSort(t *testing.T) {
    for _, tt := range testCases {
        t.Run(tt.name, func(t *testing.T) {
            actual := make([]int, len(tt.input))
            copy(actual, tt.input)

            sorting.QuickSortRandomized(actual)

            if !reflect.DeepEqual(actual, tt.expected) {
                t.Errorf("QuicksortRandomized(%v) = %v; want %v\n", tt.input, actual, tt.expected)
            }
        })
    }
}


func TestBubbleSort(t *testing.T) {
    for _, tt := range testCases {
        t.Run(tt.name, func(t *testing.T) {
            actual := make([]int, len(tt.input))
            copy(actual, tt.input)

            sorting.BubbleSort(actual)
            if !reflect.DeepEqual(actual, tt.expected) {
                t.Errorf("BubbleSort(%v) = %v; want %v\n", tt.input, actual, tt.expected)
            }
        })
    }
}
