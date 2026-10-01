func twoSum(nums []int, target int) []int {
    seen := map[int]int{nums[0]:0}

    for i, num := range nums{
        if i == 0 {continue}

        if idx, ex := seen[target - num]; ex {
            return []int{idx, i}
        }

        seen[num] = i
    }

    return []int{-1, -1}
}
