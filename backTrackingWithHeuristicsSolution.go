// This is the naive implementation of backtracking.

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
func solve(board [][]byte) bool {
    var heuristicBoard [9][9]int = assessBoard(board)
    //fmt.Print(heuristicBoard)
    var currBest = 0
    var currBestii = 0
    var currBestjj = 0
    // Depth first search.
    for ii := 0; ii < SQUARE_LENGTH; ii++ {
        for jj := 0; jj < SQUARE_LENGTH; jj++ {
            if ((currBest == 0 || heuristicBoard[ii][jj] < currBest) && heuristicBoard[ii][jj] != 0) {
                currBest = heuristicBoard[ii][jj]
                currBestii = ii
                currBestjj = jj
            }
        }
    }
    if currBest == 0 && !isSolved(board) {
        return false
    }
    if currBest == 0 {
        return true
    }
            
    // For every valid value that fits within this box, call solve on it.
    // This will resemble a decision tree.
	for val := byte('1'); val <= byte('9'); val++ {
        // Check to see if this value fits into the current solution.
		if isValid(board, currBestii, currBestjj, val) {
			board[currBestii][currBestjj] = val
			if solve(board) {
				return true
			} else {
                // We leave the '.' in this box so we can see where it went wrong.
				board[currBestii][currBestjj] = '.'
			}
        } 
	}
    // No values fit. Break this branch.
	return false
}

func prioritize(board [][]byte) []int {
    var valueQueue []int
    var positionQueue []int
    //var tempValueQueue []int
    //var tempPositionQueue []int
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
                        //tempValueQueue = append(valueQueue[:kk], valuesAllowedInCurr)
                        //tempPositionQueue = append(positionQueue[:kk], ((ii * 10) + jj))
                        //for ll := kk; ll < len(valueQueue); ll++ {
                        //    tempValueQueue = append(tempValueQueue, valueQueue[ll])
                        //}
                        //for mm := kk; mm < len(positionQueue); mm++ {
                        //    tempPositionQueue = append(tempPositionQueue, valueQueue[mm])
                        //}

                        break
                    }
                }
            }
            //valueQueue = tempValueQueue
            fmt.Print(valueQueue)
            fmt.Print(positionQueue)
            //positionQueue = tempPositionQueue
            //tempValueQueue = tempValueQueue[:0]
            //tempPositionQueue = tempPositionQueue[:0]
            valuesAllowedInCurr = 0
        }
    }
    //fmt.Print(valueQueue)
    //fmt.Print(positionQueue)
    return positionQueue
}

// Invokes a helper method that can recursively call itself then return the valid solution.
func solveSudoku(board [][]byte) {
    var heuristicQueue []int = prioritize(board)
    heuristicQueue = append(heuristicQueue, 1)
    solve(board)
}
