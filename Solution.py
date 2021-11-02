class Solution:
    # Modifies the current list of valid numbers to only include valid numbers for the row.
    def allowedInRow(possibleNumbers, board, ii, jj):
        stillPossible = ''
        for x in possibleNumbers:
            # Have not seen this number in the row by default.
            inRow = False
            for y in board[ii]:
                # If this number already exists in the row, do not add it to the output list.
                if x == y:
                    inRow = True
                    break
            if not inRow:
                stilPossible += x
                    
        # Return what is still possible after trimming down.
        return stillPossible
    
    # Modifies the current list of valid numbers to only include valid numbers for the column.
    def allowedInColumn(possibleNumbers, board, ii, jj):
        stillPossible = ''
        for x in possibleNumbers:
            # Have not seen this number in the column by default.
            inColumn = False
            for y in range(len(board)):
                # If this number already exists in the column, do not add it to the output list.
                if x == board[ii][y]:
                    inColumn = True
                    break
            if not inColumn:
                stilPossible += x
                    
        # Return what is still possible after trimming down.
        return stillPossible

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
            return possibleNumbers[0]
        else:
            return '.'
    
    # Solves the puzzle in-place.
    def solveSudoku(self, board: List[List[str]]) -> None:
        # Goes through by a Depth-First Search.
        for ii in range(len(board)):
            for jj in range(len(board[ii])):
                curr = board[ii][jj]
                if curr == '.':
                    # Check if there is only one possible option here. If so, assign the number.
                    num = checkIfOnlySolution(board, ii, jj)
                    board[ii][jj] = num
