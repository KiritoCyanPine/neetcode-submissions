class Solution:

    def encode(self, strs: List[str]) -> str:
        return "".join(chr(len(s))+s for s in strs)

    def decode(self, s: str) -> List[str]:
        words = []
        cursor = 0
        while cursor < len(s): # remember to update cursor
            lengthToRead = ord(s[cursor]) 
            offset = cursor+1+lengthToRead
            words.append(s[cursor+1: offset]) # read from next char to end
            cursor = offset
        return words