func hasDuplicate(nums []int) bool {
    duplicates := make(map[int]struct{},0)

	for _, num := range nums {
		if _, ex := duplicates[num]; ex {
			return true
		} else {
			duplicates[num] = struct{}{}
		}
	}

	return false
}
