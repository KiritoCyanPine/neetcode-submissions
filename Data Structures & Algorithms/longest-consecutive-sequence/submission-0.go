func longestConsecutive(nums []int) int {
	numSet := map[int]struct{}{}

	for _,j := range nums {
		numSet[j]= struct{}{}
	}

	curLen, maxLen := 0,0
	for num, _ := range numSet {
		if _, leftExist := numSet[num-1]; leftExist {
			continue
		}
		
		curLen = 1
		i := 1
		for {
			if _,ex := numSet[num+i];ex {
				i++
				curLen++
			} else {
				break
			}
		}
		if maxLen < curLen {
				maxLen = curLen
				curLen = 0
			}
	}
	return maxLen
}
