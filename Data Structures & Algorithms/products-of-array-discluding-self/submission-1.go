func productExceptSelf(nums []int) []int {
	// -> only 1 0
	// -> 2+ 0s 
	//
	zeroIndex := -1
	prod := 1
	ret := make([]int, len(nums), len(nums))
	for idx, num := range nums {
		if num == 0{
			// have we encountered 0 already ?
			if zeroIndex != -1 {
				return ret
			} else {
				zeroIndex = idx
			}

			continue
		}

		prod*=num
	}

	if zeroIndex != -1{
		ret[zeroIndex] = prod
		return ret
	}

	for idx, num := range nums {
		ret[idx] = prod/num
	}

	return ret
}
