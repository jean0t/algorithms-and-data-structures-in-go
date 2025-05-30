package sorting

func QuickSort(arr []int) {
    if len(arr) <= 1 {
        return
    }

    sort(arr, 0, len(arr) - 1)
}

func sort(arr []int, low, high int) {
    if low < high {
       var p int = partition(arr, low, high)

       sort(arr, low, p - 1)
       sort(arr, p+1, high)
    }
}

func partition(arr []int, low, high int) int {
    var pivot int = arr[high]
    var i int = low - 1

    for j := low; j < high; j++ {
        if arr[j] <= pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }

    arr[i+1], arr[high] = arr[high], arr[i+1]
    return i+1
}
