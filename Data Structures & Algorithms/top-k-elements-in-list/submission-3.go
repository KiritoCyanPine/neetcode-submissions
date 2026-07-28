func topKFrequent(nums []int, k int) []int {
    fmap := map[int]int{}
    for _,j := range nums {
        fmap[j]++
    }

    arr := make([][2]int, 0, len(fmap))
    for i,j := range fmap {
        arr = append(arr, [2]int{i,j})
    }

    sort.Slice(arr, func(i,j int) bool{
        return arr[i][1] > arr[j][1]
    })

    res := []int{}
    for _,j := range arr {
        res = append(res, j[0])

        k--
        if k == 0 {
            return res
        }
    }

    return res
}
