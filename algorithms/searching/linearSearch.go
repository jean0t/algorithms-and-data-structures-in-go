package searching

func LinearSearch(arr []int, want int) (int, bool) {
    var index int = -1
    var found bool = false

    for i, v := range arr {
        if v == want {
            index = i
            found = true
            break
        }
    }

    return index, found
}
