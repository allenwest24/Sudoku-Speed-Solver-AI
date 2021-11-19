package main

import (
	"fmt"
)

const SQUARE_LENGTH = 25
const ROOT = 5

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
        boxRow := ((row / ROOT) * ROOT) + (ii / ROOT)
        boxColumn := ((column / ROOT) * ROOT) + (ii % ROOT)
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
            for val := byte(1); val <= byte(25); val++ {
                // Check to see if this value fits into the current solution.
				if isValid(board, ii, jj, val) {
					valuesAllowedInCurr++
                }
            }
            if len(valueQueue) == 0 {
                valueQueue = append(valueQueue, valuesAllowedInCurr)
                positionQueue = append(positionQueue, ((ii * 100) + jj))
                continue
            } else {
                for kk := 0; kk <= len(valueQueue); kk++ {
                    if kk == len(valueQueue) {
                        valueQueue = append(valueQueue, valuesAllowedInCurr)
                        positionQueue = append(positionQueue, ((ii * 100) + jj))
                        break
                    }
                    if valueQueue[kk] >= valuesAllowedInCurr {
                        valueQueue = append(valueQueue[:kk+1], valueQueue[kk:]...)
                        valueQueue[kk] = valuesAllowedInCurr
                        positionQueue = append(positionQueue[:kk+1], positionQueue[kk:]...)
                        positionQueue[kk] = ((ii * 100) + jj)
                        break
                    }
                }
            }
            //fmt.Print(valueQueue)
            //fmt.Print(positionQueue)
            valuesAllowedInCurr = 0
        }
    }
    fmt.Print(valueQueue)
    return positionQueue
}

// Invokes a helper method that can recursively call itself then return the valid solution.
func main() {
    board := [][]byte {{11,21,byte('.'),byte('.'),byte('.'),3,byte('.'),16,byte('.'),24,13,15,byte('.'),19,byte('.'),byte('.'),byte('.'),10,byte('.'),byte('.'),byte('.'),byte('.'),4,17,byte('.')},{2,12,byte('.'),4,17,byte('.'),byte('.'),5,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),18,19,20,byte('.'),byte('.'),10,byte('.'),byte('.')},{byte('.'),byte('.'),byte('.'),10,23,byte('.'),byte('.'),7,byte('.'),17,11,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),1,16,22,byte('.'),byte('.'),byte('.'),18,byte('.'),byte('.')},{23,13,byte('.'),18,byte('.'),6,9,byte('.'),10,byte('.'),2,12,7,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),8,byte('.'),1,16,byte('.'),byte('.')},{byte('.'),3,byte('.'),byte('.'),22,byte('.'),13,15,18,byte('.'),byte('.'),9,14,10,23,byte('.'),12,byte('.'),4,byte('.'),11,21,byte('.'),byte('.'),byte('.')},{8,11,byte('.'),byte('.'),25,byte('.'),3,1,16,byte('.'),20,13,15,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),10,byte('.'),2,12,7,byte('.'),17},{byte('.'),2,12,byte('.'),byte('.'),byte('.'),byte('.'),21,5,byte('.'),byte('.'),3,1,byte('.'),22,byte('.'),byte('.'),15,18,19,6,byte('.'),14,byte('.'),byte('.')},{23,byte('.'),9,14,byte('.'),byte('.'),2,12,7,4,8,11,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),1,16,22,20,13,byte('.'),byte('.'),byte('.')},{byte('.'),20,13,byte('.'),byte('.'),byte('.'),6,9,byte('.'),byte('.'),17,byte('.'),12,byte('.'),byte('.'),8,11,byte('.'),5,25,24,byte('.'),1,byte('.'),byte('.')},{byte('.'),byte('.'),3,1,16,byte('.'),byte('.'),13,byte('.'),18,byte('.'),6,byte('.'),14,byte('.'),byte('.'),2,byte('.'),7,4,byte('.'),byte('.'),byte('.'),byte('.'),byte('.')},{byte('.'),byte('.'),byte('.'),21,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),16,19,20,byte('.'),15,18,byte('.'),byte('.'),byte('.'),byte('.'),10,17,byte('.'),byte('.'),byte('.'),4},{4,17,2,12,byte('.'),25,byte('.'),byte('.'),21,byte('.'),byte('.'),24,byte('.'),1,byte('.'),19,20,byte('.'),byte('.'),18,23,byte('.'),byte('.'),byte('.'),10},{byte('.'),byte('.'),byte('.'),9,14,4,byte('.'),byte('.'),12,byte('.'),25,8,byte('.'),byte('.'),5,22,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),13,15,byte('.')},{byte('.'),byte('.'),20,13,15,10,23,byte('.'),9,14,4,17,2,byte('.'),byte('.'),byte('.'),8,byte('.'),21,byte('.'),byte('.'),byte('.'),byte('.'),1,16},{byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),18,byte('.'),byte('.'),13,15,10,23,byte('.'),byte('.'),14,byte('.'),byte('.'),2,byte('.'),7,byte('.'),8,byte('.'),21,5},{byte('.'),25,byte('.'),11,byte('.'),16,22,byte('.'),3,byte('.'),byte('.'),19,byte('.'),byte('.'),15,10,23,byte('.'),byte('.'),byte('.'),byte('.'),17,2,12,7},{byte('.'),4,byte('.'),2,byte('.'),byte('.'),25,byte('.'),byte('.'),byte('.'),16,22,24,3,1,byte('.'),byte('.'),20,byte('.'),15,byte('.'),byte('.'),6,byte('.'),14},{byte('.'),byte('.'),23,6,byte('.'),7,byte('.'),17,2,byte('.'),byte('.'),byte('.'),8,byte('.'),21,16,byte('.'),24,3,1,18,byte('.'),20,byte('.'),15},{15,18,19,20,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),9,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),5,byte('.'),byte('.'),byte('.'),byte('.'),16,byte('.'),24,3,byte('.')},{1,byte('.'),byte('.'),24,3,15,18,19,20,13,14,10,23,byte('.'),9,7,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),25,8,byte('.'),21},{21,5,byte('.'),8,11,byte('.'),16,22,24,3,15,18,19,20,byte('.'),14,byte('.'),byte('.'),6,9,byte('.'),byte('.'),byte('.'),2,12},{byte('.'),byte('.'),4,17,byte('.'),21,5,byte('.'),byte('.'),11,1,byte('.'),22,24,byte('.'),byte('.'),18,19,20,13,byte('.'),10,byte('.'),6,byte('.')},{9,14,10,byte('.'),byte('.'),12,byte('.'),4,byte('.'),2,byte('.'),5,25,8,11,1,byte('.'),byte('.'),24,3,15,18,byte('.'),20,byte('.')},{byte('.'),15,18,byte('.'),byte('.'),9,byte('.'),10,byte('.'),6,12,byte('.'),4,17,2,21,5,byte('.'),8,11,1,byte('.'),byte('.'),24,3},{byte('.'),1,16,byte('.'),24,byte('.'),15,18,byte('.'),20,9,byte('.'),byte('.'),23,6,byte('.'),7,byte('.'),17,2,21,5,byte('.'),8,byte('.')}}
    var heuristicQueue []int = prioritize(board)
    fmt.Print(heuristicQueue)
}
