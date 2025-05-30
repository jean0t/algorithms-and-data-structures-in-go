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
    sort2(arr, 0, len(arr)-1)
}

func sort2(arr []int, low, high int) {
    if low < high {
        var p int = partition2(arr, low, high)
        
        sort2(arr, low, p-1)
        sort2(arr, p+1, high)
    }
}

func partition2(arr []int, low, high int) int {
    var randomIndex int = rand.Intn(high + 1 - low) + low
    arr[randomIndex], arr[high] = arr[high], arr[randomIndex]
    
    var pivot int = arr[high]
    var i int = low

    for j := low; j < high; j++ {
        if arr[j] < pivot {
            arr[i], arr[j] = arr[j], arr[i]
            i++
        }
    }

    arr[i], arr[high] = arr[high], arr[i]
    return i
}


