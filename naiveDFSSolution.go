// Define side size of board.
const SQUARE_LENGTH = 9

// Has the puzzle been solved or not.
func isSolved(board [][]byte) bool {
    var solved = true
    for ii := 0; ii < SQUARE_LENGTH; ii++ {
        for jj := 0; jj < SQUARE_LENGTH; jj++ {
            if board[ii][jj] == '.' {
                solved = false
                break
            }
        }
    }
    return solved
}

// Determines if a given value at a given location in a puzzle is a valid placement.
func isValid(board [][]byte, row, column int, val byte) bool {
    for ii := 0; ii < SQUARE_LENGTH; ii++ {
        // Is this value already in this column?
        if board[ii][column] == val {
            return false
        }
        // Is this value already in this row?
        if board[row][ii] == val {
	    return false
	}
        // Is this value already in this box?
        boxRow := ((row / 3) * 3) + (ii / 3)
        boxColumn := ((column / 3) * 3) + (ii % 3)
	if board[boxRow][boxColumn] == val {
	    return false
	}
    }
    return true
}

// Main method for this solution.
func solveSudoku(board [][]byte)  {
    // While this puzzle is not zolved keep solving.
    for !isSolved(board) {
	// Perform a depth-first search for the next solvable square.
        for ii := 0; ii < SQUARE_LENGTH; ii++ {
            for jj := 0; jj < SQUARE_LENGTH; jj++ {
                var allowedValues []byte
                if (board[ii][jj] != '.') {
                    continue
                }
		// Determine all values that are currently allowed in this square.
                for val := byte('1'); val <= byte('9'); val++ {
                    // Check to see if this value fits into the current solution.
                    if isValid(board, ii, jj, val) {
                        allowedValues = append(allowedValues, val)
                    }
                }
		 // If there is only one valid number for this square, we can place it and continue.
                 if len(allowedValues) == 1 {
                    board[ii][jj] = allowedValues[0]
                } else {
                    board[ii][jj] = '.'
                }
            }
        }
    }
}
