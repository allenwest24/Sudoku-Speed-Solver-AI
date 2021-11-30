package main

import (
	"fmt"
)

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
            //fmt.Print(valueQueue)
            //fmt.Print(positionQueue)
            valuesAllowedInCurr = 0
        }
    }
    return positionQueue
}

// Invokes a helper method that can recursively call itself then return the valid solution.
func main() {
    board := [][]byte {
	    {byte('.'),byte('4'),byte('.'),byte('.'),byte('.'),byte('2'),byte('.'),byte('.'),byte('.')},
	    {byte('.'),byte('.'),byte('8'),byte('4'),byte('7'),byte('.'),byte('.'),byte('.'),byte('5')},
	    {byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),byte('6'),byte('.'),byte('7'),byte('.')},
	    {byte('.'),byte('6'),byte('.'),byte('9'),byte('1'),byte('.'),byte('.'),byte('5'),byte('.')},
	    {byte('3'),byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),byte('8')},
	    {byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),byte('7'),byte('.'),byte('.'),byte('.')},
	    {byte('.'),byte('9'),byte('.'),byte('5'),byte('4'),byte('.'),byte('.'),byte('1'),byte('.')},
	    {byte('.'),byte('.'),byte('6'),byte('.'),byte('.'),byte('.'),byte('9'),byte('.'),byte('.')},
	    {byte('.'),byte('.'),byte('.'),byte('2'),byte('.'),byte('.'),byte('.'),byte('.'),byte('.')}}
    fmt.Print(prioritize(board))
    fmt.Print("\n\n")
}
