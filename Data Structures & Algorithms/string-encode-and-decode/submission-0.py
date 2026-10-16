class Solution:

    def encode(self, strs: List[str]) -> str:
        out = ""
        for s in strs:
            out += chr(len(s))+s
        return out

    def decode(self, s: str) -> List[str]:
        words = []
        cursor = 0
        while cursor < len(s): # remember to update cursor and i 
            lengthToRead = ord(s[cursor]) 
            word = s[cursor+1: cursor+1+lengthToRead]
            cursor = cursor+1+lengthToRead
            words.append(word)
        return words