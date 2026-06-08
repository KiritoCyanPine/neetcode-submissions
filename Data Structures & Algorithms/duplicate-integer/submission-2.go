func hasDuplicate(nums []int) bool {
    mp := make(map[int]bool)

    for _, number := range nums {
        if mp[number] == true {
            return true
        } 
        mp[number] = true
    }

    return false
}
