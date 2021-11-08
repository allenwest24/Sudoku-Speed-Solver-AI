const SQUARE_LENGTH = 9

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

func assessBoard(board [][]byte) [9][9]int {
    var heuristicBoard [9][9]int
    var valuesAllowedInCurr int = 0
    for ii := 0; ii < SQUARE_LENGTH; ii++ {
        for jj := 0; jj < SQUARE_LENGTH; jj++ {
            if (board[ii][jj] != '.') {
                heuristicBoard[ii][jj] = 0
				continue
			}
            for val := byte('1'); val <= byte('9'); val++ {
                // Check to see if this value fits into the current solution.
				if isValid(board, ii, jj, val) {
					valuesAllowedInCurr++
                }
            }
            heuristicBoard[ii][jj] = valuesAllowedInCurr
            valuesAllowedInCurr = 0
        }
    }
    return heuristicBoard
}

// Solve the sudoku puzzle.
func solve(board [][]byte, heuristicQueue []int) bool {
    var currii int
    var currjj int
    if len(heuristicQueue) == 0 {
        return true
    } else {
        currii = heuristicQueue[0] / 10
        currjj = heuristicQueue[0] % 10
        fmt.Print(currii, currjj)
    }
	for val := byte('1'); val <= byte('9'); val++ {
        // Check to see if this value fits into the current solution.
		if isValid(board, currii, currjj, val) {
			board[currii][currjj] = val
            if solve(board, heuristicQueue[1:]) {
				return true
			} else {
                // We leave the '.' in this box so we can see where it went wrong.
				board[currii][currjj] = '.'
			}
        } 
	}
    // No values fit. Break this branch.
	return false
}

func prioritize(board [][]byte) []int {
    var valueQueue []int
    var positionQueue []int
    var valuesAllowedInCurr int = 0
    for ii := 0; ii < SQUARE_LENGTH; ii++ {
        for jj := 0; jj < SQUARE_LENGTH; jj++ {
            if (board[ii][jj] != '.') {
				continue
			}
            for val := byte('1'); val <= byte('9'); val++ {
                // Check to see if this value fits into the current solution.
				if isValid(board, ii, jj, val) {
					valuesAllowedInCurr++
                }
            }
            if len(valueQueue) == 0 {
                valueQueue = append(valueQueue, valuesAllowedInCurr)
                positionQueue = append(positionQueue, ((ii * 10) + jj))
                continue
            } else {
                for kk := 0; kk <= len(valueQueue); kk++ {
                    if kk == len(valueQueue) {
                        valueQueue = append(valueQueue, valuesAllowedInCurr)
                        positionQueue = append(positionQueue, ((ii * 10) + jj))
                        break
                    }
                    if valueQueue[kk] >= valuesAllowedInCurr {
                        valueQueue = append(valueQueue[:kk+1], valueQueue[kk:]...)
                        valueQueue[kk] = valuesAllowedInCurr
                        positionQueue = append(positionQueue[:kk+1], positionQueue[kk:]...)
                        positionQueue[kk] = ((ii * 10) + jj)
                        break
                    }
                }
            }
            fmt.Print(valueQueue)
            fmt.Print(positionQueue)
            valuesAllowedInCurr = 0
        }
    }
    return positionQueue
}

// Invokes a helper method that can recursively call itself then return the valid solution.
func solveSudoku(board [][]byte) {
    var heuristicQueue []int = prioritize(board)
    solve(board, heuristicQueue)
}
