class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        map = {}

        for i, val in enumerate(nums):
            if target-val in map.keys():
                return [map[target-val], i]
            map[val] = i
        return [-1,-1]
