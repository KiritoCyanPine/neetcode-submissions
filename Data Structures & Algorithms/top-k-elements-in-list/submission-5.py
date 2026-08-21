import heapq
class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        min_heap = []

        fm = self.countFreq( nums)
        for key, freq in fm.items():
            heapq.heappush(min_heap, (freq, key))

            if len(min_heap) > k:
                heapq.heappop(min_heap)

        return [num for _, num in min_heap]
    
    def countFreq(self, nums: List[int]) -> Dict:
        fm = {}
        for num in nums:
            fm[num] = fm.get(num,0)+1
        return fm
        