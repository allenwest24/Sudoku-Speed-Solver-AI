package main

import (
        "fmt"
	"strconv"
	"time"
)

// Size of side of square.
const SQUARE_LENGTH = 9

// Checks if a given value is valid at a given position.
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

// Solve the sudoku puzzle.
func solve(board [][]byte, branch []int) bool {
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
		    formattedVal, _ := strconv.Atoi(string(val))
		    branch = append(branch, formattedVal)
		    if solve(board, branch) {
		        return true
		    } else {
                        // We leave the '.' in this box so we can see where it went wrong.
			board[ii][jj] = '.'
		    }
                } 
	    }
            // No values fit. Break this branch.
	    fmt.Print(branch)
	    fmt.Print("\n\n")
	    time.Sleep(2 * time.Second)
	    return false
	}
    }
    // If we got through the entire board, we have successfully solved this puzzle.
    // Backtrack to the root of the tree and say so.
    fmt.Print(branch)
    fmt.Print("\n\n")
    return true
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
    branch := []int {}
    solve(board, branch)
}
