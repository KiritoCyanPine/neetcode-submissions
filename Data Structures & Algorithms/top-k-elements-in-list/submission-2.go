func topKFrequent(nums []int, k int) []int {
    fmap := map[int]int{}

    for _,j := range nums {
        fmap[j]++
    }

    freqToNum := make([][]int, len(nums)+1, len(nums)+1)

    for num,freq := range fmap {
        freqToNum[freq] = append(freqToNum[freq], num)
    }

    res := []int{}

        for i:=len(freqToNum)-1;i>=0;i--{
            if len(freqToNum[i]) == 0{
                continue
            }

            for j:=0; j<=len(freqToNum[i])-1;j++ {
                res = append(res, freqToNum[i][j])

                k--

                if k == 0{
                    return res
                }
            }
        }

    return res
}
