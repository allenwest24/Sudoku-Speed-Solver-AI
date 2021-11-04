// This is the naive implementation of backtracking.

const SQUARE_LENGTH = 9

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

func solveSudoku(board [][]byte) {
	// Depth first search.
    for ii := 0; ii < SQUARE_LENGTH; ii++ {
        for jj := 0; jj < SQUARE_LENGTH; jj++ {
            var allowedValues = []byte
            if (board[ii][jj] != '.') {
				continue
			}
			for val := byte('1'); val <= byte('9'); val++ {
                // Check to see if this value fits into the current solution.
				if isValid(board, ii, jj, val) {
                    allowedValues = append(allowedValues, val)
			}
            if len(allowedValues) == 1 {
                board[ii][jj] = allowedValues[0]
            } else {
                board[ii][jj] = '.'
            }
		}
	}
}

