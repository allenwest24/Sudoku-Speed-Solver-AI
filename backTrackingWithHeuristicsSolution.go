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
    // If this value is not found in any conflicting locations, it is valid.
	return true
}

func assessBoard(board [][]byte) [][]int {
    var heuristicBoard [][]int
    var valuesAllowedInCurr int = 0
    var tempRow []int
    for ii := 0; ii < SQUARE_LENGTH; ii++ {
        for jj := 0; jj < SQUARE_LENGTH; jj++ {
            if (board[ii][jj] != '.') {
                tempRow = append(tempRow, valuesAllowedInCurr)
                valuesAllowedInCurr = 0
				continue
			}
            for val := byte('1'); val <= byte('9'); val++ {
                // Check to see if this value fits into the current solution.
				if isValid(board, ii, jj, val) {
					valuesAllowedInCurr++
                }
            }
            tempRow = append(tempRow, valuesAllowedInCurr)
            valuesAllowedInCurr = 0
        }
        heuristicBoard = append(heuristicBoard, tempRow)
        tempRow = tempRow[:0]
    }
    return heuristicBoard
}

// Solve the sudoku puzzle.
func solve(board [][]byte) bool {
    var heuristicBoard [][]int = assessBoard(board)
    fmt.Print(heuristicBoard)
    // Depth first search.
    for ii := 0; ii < SQUARE_LENGTH; ii++ {
        for jj := 0; jj < SQUARE_LENGTH; jj++ {
            // If this square has already been solved, move down the branch.
            if (board[ii][jj] != '.') {
				continue
			}
            
            // For every valid value that fits within this box, call solve on it.
            // This will resemble a decision tree.
			for val := byte('1'); val <= byte('9'); val++ {
                // Check to see if this value fits into the current solution.
				if isValid(board, ii, jj, val) {
					board[ii][jj] = val
					if solve(board) {
						return true
					} else {
                        // We leave the '.' in this box so we can see where it went wrong.
						board[ii][jj] = '.'
					}
                } 
			}
            // No values fit. Break this branch.
			return false
		}
	}
    // If we got through the entire board, we have successfully solved this puzzle.
    // Backtrack to the root of the tree and say so.
	return true
}

// Invokes a helper method that can recursively call itself then return the valid solution.
func solveSudoku(board [][]byte) {
	solve(board)
}

