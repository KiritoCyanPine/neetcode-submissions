func twoSum(nums []int, target int) []int {
    seen := map[int]int{}

    for i, num := range nums{

        if idx, ex := seen[target - num]; ex {
            return []int{idx, i}
        }

        seen[num] = i
    }

    return []int{-1, -1}
}
