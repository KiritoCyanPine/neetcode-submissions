
func topKFrequent(nums []int, k int) []int {
    fmap := map[int]int{}
    for _,num := range nums {
        fmap[num]++
    }

    repeatitionMap := make([][]int, len(nums), len(nums))
    for num,freq := range fmap {
        repeatitionMap[freq-1] = append(repeatitionMap[freq-1], num)
    }

    kfreq := []int{}
    for i:=len(repeatitionMap)-1; i>=0;i--{
        for _,num := range repeatitionMap[i]{
            kfreq = append(kfreq, num)
            k--
            if k == 0 {
                return kfreq
            }
        }
        if k == 0 {
                return kfreq
            }
    }
    return kfreq
}
