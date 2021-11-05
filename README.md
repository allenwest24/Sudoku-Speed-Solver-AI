# Sudoku-Speed-Solver-AI

## Solutions:
 - naiveDFSSolution: Goes through a depth-first dearch of the board and only places values if they are the only possible value for this square.
 - naiveBacktrackingDFSSolution: Goes through depth-first dearch of the board and uses a decision tree to go down a branch for every possible value until one results in the puzzle being solved.

# Tests:
tests.txt is a file that contains 3 easy, 3 medium, 3 hard, 3 expert, and 3 evil puzzles. 

# Results:
Upon running runner.go in a command line, it will run all of the different solutions and time them very accurately. It will notify how hard of a puzzle they could solve, and what the average time for each difficulty it took for each solution, allowing us to compare the results.
