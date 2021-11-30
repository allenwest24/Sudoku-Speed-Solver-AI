package main

import (
        "fmt"
	"time"
)

const SQUARE_LENGTH = 9
const ROOT = 3

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
        boxRow := ((row / ROOT) * ROOT) + (ii / ROOT)
        boxColumn := ((column / ROOT) * ROOT) + (ii % ROOT)
		if board[boxRow][boxColumn] == val {
			return false
		}
	}
    // If this value is not found in any conflicting locations, it is valid.
	return true
}

// Solve the sudoku puzzle.
func solve(board [][]byte, heuristicQueue []int) bool {
    var currii int
    var currjj int
    if len(heuristicQueue) == 0 {
	solved := isSolved(board)
        return solved
    } else {
        currii = heuristicQueue[0] / 10
        currjj = heuristicQueue[0] % 10
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
    var heuristicQueue = []int{43, 74, 65, 45, 75, 62, 60, 53, 44, 35, 32, 23, 15, 11, 3, 88, 87, 85, 84, 78, 77, 73, 68, 51, 47, 41, 38, 36, 30, 24, 21, 20, 17, 16, 10, 8, 7, 6, 4, 82, 81, 80, 66, 57, 56, 54, 52, 46, 28, 26, 22, 0, 86, 71, 70, 58, 50, 42, 2}
    start := time.Now()
    solve(board, heuristicQueue)
    t := time.Now()
    elapsed := t.Sub(start)
    fmt.Print(elapsed)
    fmt.Print("\n\n")
}
