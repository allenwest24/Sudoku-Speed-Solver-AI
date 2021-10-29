class Solution:
    # TODO: Implement.
    # Modifies the current list of valid numbers to only include valid numbers for the Row.
    def allowedInRow(possibleNumbers, board, ii, jj):
        return possibleNumbers
    
    # TODO: Implement.
    # Modifies the current list of valid numbers to only include valid numbers for the Column.
    def allowedInColumn(possibleNumbers, board, ii, jj):
        return possibleNumbers

    # TODO: Implement.
    # Modifies the current list of valid numbers to only include valid numbers for the box.
    def allowedInBox(possibleNumbers, board, ii, jj):
        return possibleNumbers
    
    # Sees if there is only one possible solution, making it a definite placement.
    def checkIfOnlySolution(board, ii, jj):
        possibleNumbers = '123456789'
        
        # TODO: Add more heuristics to narrow down the options.
        possibleNumbers = allowedInRow(possibleNumbers, board, ii, jj)
        possibleNumbers = allowedInColumn(possibleNumbers, board, ii, jj)
        possibleNumbers = allowedInBox(possibleNumbers, board, ii, jj)
        if len(possibleNumbers) == 1:
            return True, possibleNumbers[0]
    
    # Solves the puzzle in-place.
    def solveSudoku(self, board: List[List[str]]) -> None:
        # Goes through by a Depth-First Search.
        for ii in range(len(board)):
            for jj in range(len(board[ii])):
                curr = board[ii][jj]
                if curr == '.':
                    # Check if there is only one possible option here. If so, assign the number.
                    huh, num = checkIfOnlySolution(board, ii, jj)
                    if huh:
                        board[ii][jj] = num
                        
