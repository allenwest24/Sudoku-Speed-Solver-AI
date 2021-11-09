# Sudoku-Speed-Solver-AI

## Solutions:
The following are rough times for "evil' sudoku puzzles:
 - naiveDFSSolution: Does not solve reliably.
 - naiveBacktrackingDFSSolution: ~350 ms
 - backTrackingWithHeuristicsSolution: ~30 ms
 - backTrackingWithHeuristicsAndPreComputingSolution: ~ 5 ms

# Tests:
tests.txt is a file that contains 3 easy, 3 medium, 3 hard, 3 expert, and 3 evil puzzles. 

# Results:
Upon running runner.go in a command line, it will run all of the different solutions and time them very accurately. It will notify how hard of a puzzle they could solve, and what the average time for each difficulty it took for each solution, allowing us to compare the results.
