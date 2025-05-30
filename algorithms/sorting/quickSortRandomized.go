package sorting

import (
    "time"
    "math/rand"
)

func QuickSortRandomized(arr []int) {
    if len(arr) < 2 {
        return
    }

    rand.Seed(time.Now().UnixNano())
    sort(arr, 0, len(arr)-1)
}

func sort(arr []int, low, high int) {
    if low < high {
        var p int = partition(arr, low, high)
        
        sort(arr, low, p-1)
        sort(arr, p+1, high)
    }
}

func partition(arr []int, low, high int) int {
    var randomIndex int = rand.Intn(high + 1 - low) + low
    arr[randomIndex], arr[high] = arr[high], arr[randomIndex]
    
    var pivot int = arr[high]
    var i int = low

    for j := low; j < high; j++ {
        if arr[j] <= pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }

    arr[i], arr[high] = arr[high], arr[i]
    return i
}


