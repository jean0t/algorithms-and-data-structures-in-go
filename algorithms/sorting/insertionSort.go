package main

import "fmt"

func insertionSort(numbers []int) []int {
	var length = len(numbers)
	for i := 1; i < length; i++ {
		var j = i - 1
		for j >= 0 && numbers[j] > numbers[j+1] {
			var temp = numbers[j]
			numbers[j] = numbers[j+1]
			numbers[j+1] = temp
			j--
		}
	}
	return numbers
}

func main() {
	var numbers = []int{4, 2, 5, 1, 7}
	fmt.Println("Unsorted: ", numbers)
	var sorted = insertionSort(numbers)
	fmt.Println("Sorted: ", sorted)
}
