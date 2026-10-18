class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        # rows 
        for row in board:
            if self.validateLine(row) != True:
                return False
        
        # cols
        transpose = list(map(list, zip(*board)))
        for row in transpose:
            if self.validateLine(row) != True:
                return False
        
        # box
        for rs in range(0,9,3):
            for cs in range(0,9,3):
                l = list()

                for i in range(cs, cs+3):
                    l = l + board[i][rs:rs+3]
                if self.validateLine(l) != True:
                    return False

        return True

    
    def validateLine(self, line: List[str] ) -> bool:
        s = set()
        for item in line:
            if item == ".":
                continue
            if item in s:
                return False
            s.add(item)
        return True

    
