class Solution:
    def solveSudoku(self, board: List[List[str]]) -> None:
        def isSolved(board):
            solved = True
            for ii in range(9):
                for jj in range(9):
                    if board[ii][jj] == '.':
                        solved = False
                        break
            return solved
        
        def isValid(board, row, column, val):
            for ii in range(9):
                if board[ii][column] == val:
                    return False
                if board[row][ii] == val:
                    return False
                boxRow = int(((row / 3) * 3) + (ii / 3))
                boxColumn = int(((column / 3) * 3) + (1 % 3))
                if board[boxRow][boxColumn] == val:
                    return False
            return True
        
        for ii in range(9):
            for jj in range(9):
                allowedValues = []
                if board[ii][jj] != '.':
                    continue
                for val in range(9):
                    if isValid(board, ii, jj, str(val)):
                        allowedValues.append(str(val))
                if len(allowedValues) == 1:
                    board[ii][jj] = allowedValues[0]
                else:
                    board[ii][jj] = '.'
        
