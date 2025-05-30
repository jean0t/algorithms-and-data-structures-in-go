package tests

import (
    "testing"
    "reflect"
    "github.com/jean0t/algorithms-and-data-structures-in-go/algorithms/sorting"

)

func TestQuickSort(t *testing.T) {
    tests := []struct {
        name string
        input []int
        expected []int
    }{
        {
            name: "Empty array",
            input: []int{},
            expected: []int{},
        },
        {
            name: "Single element array",
            input: []int{5},
            expected: []int{5},
        },
        {
            name: "Already sorted array",
            input: []int{1,2,3,4,5},
            expected: []int{1,2,3,4,5},
        },
        {
            name: "Reverse sorted array",
            input: []int{5,4,3,2,1},
            expected: []int{1,2,3,4,5},
        },
        {
            name: "Array with duplicates",
            input: []int{3,3,2,7,7,1,2},
            expected: []int{1,2,2,3,3,7,7},
        },
        {
            name: "Shuffled array",
            input: []int{5,2,7,1,9,3},
            expected: []int{1,2,3,5,7,9},
        },
    }

    for _, tt := range tests {
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
    tests := []struct {
        name string
        input []int
        expected []int
    }{
        {
            name: "Empty array",
            input: []int{},
            expected: []int{},
        },
        {
            name: "Single element array",
            input: []int{5},
            expected: []int{5},
        },
        {
            name: "Already sorted array",
            input: []int{1,2,3,4,5},
            expected: []int{1,2,3,4,5},
        },
        {
            name: "Reverse sorted array",
            input: []int{5,4,3,2,1},
            expected: []int{1,2,3,4,5},
        },
        {
            name: "Array with duplicates",
            input: []int{3,3,2,7,7,1,2},
            expected: []int{1,2,2,3,3,7,7},
        },
        {
            name: "Shuffled array",
            input: []int{5,2,7,1,9,3},
            expected: []int{1,2,3,5,7,9},
        },
    }

    for _, tt := range tests {
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
