from collections import defaultdict

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        group = defaultdict(list)

        ord_a = ord('a')
        
        for string in strs:
            count = [0]*26

            for character in string:
                count[ord(character) - ord_a] += 1
            
            group[tuple(count)].append(string)
        
        return list(group.values())
