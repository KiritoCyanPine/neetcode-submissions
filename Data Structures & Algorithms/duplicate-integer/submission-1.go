func hasDuplicate(nums []int) bool {
    mp := make(map[int]struct{})

    for _, number := range nums {
        _, found := mp[number]
        if found {
            return true
        } else {
            mp[number] = struct{}{}
        }
    }

    return false
}
