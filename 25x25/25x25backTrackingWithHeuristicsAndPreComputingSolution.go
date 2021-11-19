package main

import (
        "fmt"
)

const SQUARE_LENGTH = 25
const ROOT = 5

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
    if solved {
	for kk := 0; kk < len(board); kk++ {
	    for ll := 0; ll < SQUARE_LENGTH; ll++ {
		if  kk == 0 && ll == 0 {
			fmt.Print("\n\nSolved:\n")
	        }
		fmt.Print(string(board[kk][ll]))
                fmt.Print(" | ")
                if ll == (SQUARE_LENGTH - 1) {
		    fmt.Print("\n------------------------------------\n") 
	    	}
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
    tally := 0
    for tmp := 0; tmp < 25; tmp++ {
	    for tmp2 := 0; tmp2 < 25; tmp2++ {
		    if board[tmp][tmp2] == '.' {
			    tally++
		    }
	    }
    }
    fmt.Print(tally)
    if len(heuristicQueue) == 0 {
	solved := isSolved(board)
        return solved
    } else {
        currii = heuristicQueue[0] / 100
        currjj = heuristicQueue[0] % 100
	//fmt.Print(board)
	//fmt.Print("\n")
    }
	for val := byte(1); val <= byte(SQUARE_LENGTH); val++ {
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
    board := [][]byte {{11,21,byte('.'),byte('.'),byte('.'),3,byte('.'),16,byte('.'),24,13,15,byte('.'),19,byte('.'),byte('.'),byte('.'),10,byte('.'),byte('.'),byte('.'),byte('.'),4,17,byte('.')},{2,12,byte('.'),4,17,byte('.'),byte('.'),5,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),18,19,20,byte('.'),byte('.'),10,byte('.'),byte('.')},{byte('.'),byte('.'),byte('.'),10,23,byte('.'),byte('.'),7,byte('.'),17,11,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),1,16,22,byte('.'),byte('.'),byte('.'),18,byte('.'),byte('.')},{23,13,byte('.'),18,byte('.'),6,9,byte('.'),10,byte('.'),2,12,7,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),8,byte('.'),1,16,byte('.'),byte('.')},{byte('.'),3,byte('.'),byte('.'),22,byte('.'),13,15,18,byte('.'),byte('.'),9,14,10,23,byte('.'),12,byte('.'),4,byte('.'),11,21,byte('.'),byte('.'),byte('.')},{8,11,byte('.'),byte('.'),25,byte('.'),3,1,16,byte('.'),20,13,15,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),10,byte('.'),2,12,7,byte('.'),17},{byte('.'),2,12,byte('.'),byte('.'),byte('.'),byte('.'),21,5,byte('.'),byte('.'),3,1,byte('.'),22,byte('.'),byte('.'),15,18,19,6,byte('.'),14,byte('.'),byte('.')},{23,byte('.'),9,14,byte('.'),byte('.'),2,12,7,4,8,11,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),1,16,22,20,13,byte('.'),byte('.'),byte('.')},{byte('.'),20,13,byte('.'),byte('.'),byte('.'),6,9,byte('.'),byte('.'),17,byte('.'),12,byte('.'),byte('.'),8,11,byte('.'),5,25,24,byte('.'),1,byte('.'),byte('.')},{byte('.'),byte('.'),3,1,16,byte('.'),byte('.'),13,byte('.'),18,byte('.'),6,byte('.'),14,byte('.'),byte('.'),2,byte('.'),7,4,byte('.'),byte('.'),byte('.'),byte('.'),byte('.')},{byte('.'),byte('.'),byte('.'),21,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),16,19,20,byte('.'),15,18,byte('.'),byte('.'),byte('.'),byte('.'),10,17,byte('.'),byte('.'),byte('.'),4},{4,17,2,12,byte('.'),25,byte('.'),byte('.'),21,byte('.'),byte('.'),24,byte('.'),1,byte('.'),19,20,byte('.'),byte('.'),18,23,byte('.'),byte('.'),byte('.'),10},{byte('.'),byte('.'),byte('.'),9,14,4,byte('.'),byte('.'),12,byte('.'),25,8,byte('.'),byte('.'),5,22,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),13,15,byte('.')},{byte('.'),byte('.'),20,13,15,10,23,byte('.'),9,14,4,17,2,byte('.'),byte('.'),byte('.'),8,byte('.'),21,byte('.'),byte('.'),byte('.'),byte('.'),1,16},{byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),18,byte('.'),byte('.'),13,15,10,23,byte('.'),byte('.'),14,byte('.'),byte('.'),2,byte('.'),7,byte('.'),8,byte('.'),21,5},{byte('.'),25,byte('.'),11,byte('.'),16,22,byte('.'),3,byte('.'),byte('.'),19,byte('.'),byte('.'),15,10,23,byte('.'),byte('.'),byte('.'),byte('.'),17,2,12,7},{byte('.'),4,byte('.'),2,byte('.'),byte('.'),25,byte('.'),byte('.'),byte('.'),16,22,24,3,1,byte('.'),byte('.'),20,byte('.'),15,byte('.'),byte('.'),6,byte('.'),14},{byte('.'),byte('.'),23,6,byte('.'),7,byte('.'),17,2,byte('.'),byte('.'),byte('.'),8,byte('.'),21,16,byte('.'),24,3,1,18,byte('.'),20,byte('.'),15},{15,18,19,20,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),9,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),5,byte('.'),byte('.'),byte('.'),byte('.'),16,byte('.'),24,3,byte('.')},{1,byte('.'),byte('.'),24,3,15,18,19,20,13,14,10,23,byte('.'),9,7,byte('.'),byte('.'),byte('.'),byte('.'),byte('.'),25,8,byte('.'),21},{21,5,byte('.'),8,11,byte('.'),16,22,24,3,15,18,19,20,byte('.'),14,byte('.'),byte('.'),6,9,byte('.'),byte('.'),byte('.'),2,12},{byte('.'),byte('.'),4,17,byte('.'),21,5,byte('.'),byte('.'),11,1,byte('.'),22,24,byte('.'),byte('.'),18,19,20,13,byte('.'),10,byte('.'),6,byte('.')},{9,14,10,byte('.'),byte('.'),12,byte('.'),4,byte('.'),2,byte('.'),5,25,8,11,1,byte('.'),byte('.'),24,3,15,18,byte('.'),20,byte('.')},{byte('.'),15,18,byte('.'),byte('.'),9,byte('.'),10,byte('.'),6,12,byte('.'),4,17,2,21,5,byte('.'),8,11,1,byte('.'),byte('.'),24,3},{byte('.'),1,16,byte('.'),24,byte('.'),15,18,byte('.'),20,9,byte('.'),byte('.'),23,6,byte('.'),7,byte('.'),17,2,21,5,byte('.'),8,byte('.')}}    
    fmt.Print(board)
    for ii := 0; ii < len(board); ii++ {
	    fmt.Print(len(board[ii]))
    }
    var heuristicQueue = []int{2411, 1710, 603, 2217, 2216, 2210, 2101, 2002, 1901, 1810, 1711, 2415, 2412, 2408, 2405, 2317, 2311, 2306, 2208, 2206, 2124, 2122, 2114, 2111, 2104, 2022, 2021, 2017, 2016, 2014, 1920, 1919, 1918, 1902, 1709, 1701, 1608, 1607, 1519, 1518, 1510, 1500, 1404, 1314, 1220, 1218, 1209, 1106, 724, 705, 610, 604, 514, 509, 503, 422, 410, 304, 211, 14, 2422, 2417, 2403, 2322, 2308, 2304, 2224, 2222, 2204, 2203, 2120, 2115, 2108, 2107, 2100, 2020, 2005, 1923, 1916, 1821, 1812, 1811, 1621, 1618, 1605, 1520, 1509, 1507, 1502, 1319, 1307, 1122, 1114, 1110, 1109, 1107, 1104, 914, 912, 910, 901, 824, 817, 809, 805, 803, 715, 714, 713, 712, 701, 616, 609, 600, 515, 423, 419, 409, 318, 309, 307, 112, 111, 19, 16, 12, 2424, 2303, 1824, 1819, 1807, 1805, 1721, 1706, 1704, 1620, 1616, 1602, 1513, 1512, 1504, 1418, 1412, 1407, 1403, 1401, 1313, 1301, 1219, 1217, 1216, 1213, 1212, 1202, 1123, 1118, 1012, 1008, 1006, 1004, 917, 814, 811, 808, 800, 723, 716, 624, 621, 613, 519, 517, 516, 513, 505, 502, 417, 415, 405, 403, 324, 323, 313, 302, 219, 213, 212, 208, 205, 201, 114, 113, 110, 106, 18, 15, 6, 2400, 2321, 2300, 1917, 1913, 1818, 1814, 1804, 1713, 1700, 1615, 1609, 1600, 1420, 1413, 1406, 1402, 1322, 1320, 1317, 1315, 1207, 1112, 1023, 1022, 1007, 1001, 921, 915, 900, 813, 804, 722, 704, 615, 606, 605, 320, 315, 314, 223, 214, 109, 105, 8, 1808, 1716, 1623, 1517, 1422, 1400, 1321, 1224, 1206, 1201, 1121, 1117, 1016, 1005, 1000, 924, 920, 906, 823, 821, 623, 523, 424, 402, 317, 316, 206, 202, 200, 123, 108, 4, 2, 1816, 1813, 1806, 1723, 1604, 1416, 1221, 1200, 1018, 923, 908, 905, 400, 215, 124, 120, 102, 24, 21, 20, 1415, 1300, 1021, 1015, 1002, 922, 224, 221, 220, 121, 116, 115, 1017, 3, 1817}
    solve(board, heuristicQueue)
}
