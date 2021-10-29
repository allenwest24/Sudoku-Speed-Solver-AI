class Solution:
    def allowedInRow(possibleNumbers, board, ii, jj):
        return possibleNumbers
    
    def allowedInColumn(possibleNumbers, board, ii, jj):
        return possibleNumbers
    
    def allowedInBox(possibleNumbers, board, ii, jj):
        return possibleNumbers
    
    def checkIfOnlySolution(board, ii, jj):
        possibleNumbers = '123456789'
        possibleNumbers = allowedInRow(possibleNumbers, board, ii, jj)
        possibleNumbers = allowedInColumn(possibleNumbers, board, ii, jj)
        possibleNumbers = allowedInBox(possibleNumbers, board, ii, jj)
        if len(possibleNumbers) == 1:
            return True, possibleNumbers[0]
    
    def solveSudoku(self, board: List[List[str]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        for ii in range(len(board)):
            for jj in range(len(board[ii])):
                curr = board[ii][jj]
                if curr == '.':
                    # Check if there is only one possible option here. If so, assign the number.
                    huh, num = checkIfOnlySolution(board, ii, jj)
                    if huh:
                        board[ii][jj] = num
                        
