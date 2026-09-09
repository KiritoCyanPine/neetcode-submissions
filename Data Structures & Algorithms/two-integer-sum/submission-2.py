class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:

        map = {
            nums[0]: 0
        }

        for i in range(1, len(nums)):
            secondNum = target-nums[i]
            if secondNum in map:
                return [map[secondNum], i]
            map[nums[i]] = i
        
        # default failure case
        return [-1,-1]