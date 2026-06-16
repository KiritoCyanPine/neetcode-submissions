
/*
Using the BUcket sort method to solve this

arr := [7,3,3,9,9,9,4]

====== STEP : 1 ======
All the bUckets are created based on the freqUency at which they are repeated

{
    7 = 1 times
    3 = 2 times
    9 = 3 times
    4 = 1 times
}

====== STEP : 2 ======
BUckets are created for the length of total nUms we have + 1.

U.0 U.1 U.2 U.3 U.4 U.5 U.6 U.7 

The bUckets are filled by the index of freqUency
 to the elements representing this freqUency in map

    7,4  3   9 
U.0 U.1 U.2 U.3 U.4 U.5 U.6 U.7 

====== STEP : 3 ======
Now that you have filled the buckets 
parse the buckets from the end to start. And push the elements into the result
Array one by one until the length of the result is k. 

Example if K = 3 , possible length of result is 3

      <---<----<--------- parse until we get k results
    7,4`  3`   9` 
U.0 U.1 U.2 U.3 U.4 U.5 U.6 U.7 

RES := [9,3,4]
*/

func topKFrequent(nums []int, k int) []int {
    // create a frequency map for the elements we have
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
