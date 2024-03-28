package main

import "fmt"

func binarySearch(array []int, wanted int) int {
  low := 0
  high := len(array) - 1

  for low <= high {
    middle := (low + high) / 2

    if array[middle] == wanted {
      return middle
    } else if array[middle] < wanted {
      low = middle + 1
    } else {
      high = middle - 1
    }
  }

  return -1
}


func confering(index int) {
  if index != -1 {
    fmt.Printf("Element found at index %d\n", index)
  } else {
    fmt.Printf("Element not found in the array\n")
  }
}

func main() {

  arr := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14}
  target1 := 8
  target2 := 15

  confering(binarySearch(arr, target1))
  confering(binarySearch(arr, target2))

}
