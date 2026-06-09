func twoSum(nums []int, target int) []int {
    seen := map[int]int{}

    for i, num := range nums {
        second := target-num

        secondIdx , found := seen[second]
        if found {
            return []int{secondIdx,i}
        }

        seen[num] = i
    }

    return []int{-1,-1}
}
