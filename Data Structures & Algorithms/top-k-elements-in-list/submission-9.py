import heapq
class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        minheap = []

        fm = {}
        for num in nums:
            fm[num] = fm.get(num, 0)+1
        
        for key, f in fm.items():
            heapq.heappush(minheap, (f, key))

            if len(minheap) > k:
                heapq.heappop(minheap)
        
        return [num for _, num in minheap]