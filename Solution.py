class Solution:    
    def solveSudoku(self, board: List[List[str]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
            # Take the list of currently possible numbers and trim down to only those allowed by row rules.
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
                    stillPossible += x

            # Return what is still possible after trimming down.
            return stillPossible

         # Take the list of currently possible numbers and trim down to only those allowed by column rules.
        def allowedInColumn(possibleNumbers, board, ii, jj):
            stillPossible = ''
            for x in possibleNumbers:
                # Have not seen this number in the column by default.
                inColumn = False
                for y in range(len(board)):
                    # If this number already exists in the column, do not add it to the output list.
                    if x == board[y][jj]:
                        inColumn = True
                        break
                if not inColumn:
                    stillPossible += x

            # Return what is still possible after trimming down.
            return stillPossible

        def allowedInBox(possibleNumbers, board, ii, jj):
            return possibleNumbers

        def checkIfOnlySolution(board, ii, jj):
            possibleNumbers = '123456789'
            print(possibleNumbers)
            possibleNumbers = allowedInRow(possibleNumbers, board, ii, jj)
            print(possibleNumbers)
            possibleNumbers = allowedInColumn(possibleNumbers, board, ii, jj)
            print(possibleNumbers)
            possibleNumbers = allowedInBox(possibleNumbers, board, ii, jj)
            if len(possibleNumbers) == 1:
                return possibleNumbers[0]
            else:
                print(len(possibleNumbers))
                return '.'

        for ii in range(len(board)):
            for jj in range(len(board[ii])):
                curr = board[ii][jj]
                if curr == '.':
                    # Check if there is only one possible option here. If so, assign the number.
                    num = checkIfOnlySolution(board, ii, jj)
                    board[ii][jj] = num
                        
