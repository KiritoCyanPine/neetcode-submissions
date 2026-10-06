
func topKFrequent(nums []int, k int) []int {
    fmap := map[int]int{}
    for _,num := range nums {
        fmap[num]++
    }

    // 0 : freq, 1 : num
    freqToNumArr := make([][2]int, 0, len(fmap))
    for num, freq := range fmap{
        freqToNumArr = append(freqToNumArr, [2]int{freq, num})
    }


    sort.Slice(freqToNumArr, func(i,j int) bool{
        return freqToNumArr[i][0] > freqToNumArr[j][0]
    })


    res := make([]int, k, k)
    for i:=0; i<k;i++ {
        res[i] = freqToNumArr[i][1]
    }

    return res
}
