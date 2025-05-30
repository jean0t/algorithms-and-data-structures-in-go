package sorting

import "fmt"

func bubbleSort(numbers []int) []int {
	var length = len(numbers) - 1
	for i := 0; i <= length; i++ {
		for j := 0; j <= length-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}
