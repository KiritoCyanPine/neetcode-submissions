
/*


U. u. u. u. 


*/

func topKFrequent(nums []int, k int) []int {
    counts := map[int]int{}
    for _,j:=range nums{
        counts[j]++
    }

    buckets := make([][]int, len(nums)+1, len(nums)+1)
    for number, freq := range counts {
        buckets[freq] = append(buckets[freq], number)
    }


    result := []int{}
    for i:= len(buckets)-1 ; i>=0;i--{
        if len(buckets[i]) == 0 {
            continue
        }

        for k>0 && len(buckets[i]) > 0 {
            last := buckets[i][len(buckets[i])-1]
            buckets[i] = buckets[i][:len(buckets[i])-1]
            result = append(result, last)
            k--
        }
    }

    return result
}
