# Speedy Sudoku AI
By Allen West and
Walter Geanacopoulos

## Guide to Contents of this Repository:
- demo/: this contains both live demos, fast and slow, as well as the three solutions pre-loaded with an evil puzzle that will print out the time it takes to solve it.
- tools/: homemade tools we had to make in order to format queues, boards, and other things.
- naiveDFSSolution.go: SLOW  - uses depth-first search and basic rules to solve only if one value is valid in a square.
- naiveBackTrackingDFSSolution.go: MORE RELIABLE - first attempt at using the backtracking algorithm.
- backTrackingWithHeuristicsAndPreComputingSolution.go: FAST - uses heuristics to determine which squares we backtrack on first.
- backTrackingWithOrderedVals.go: INSANELY FAST - combines all previous enhancements with careful ordering of the values we try to place first.
To see how these solutions compared to each other in the end please go to our results section below!

## Introduction:
To understand what we wanted to accomplish we first have to understand what Sodoku is and how the game is played. A typical game board is of equal length to equal height. Even though there may be other boards and combinations, we stuck to the standard boards mentioned above. At the start of the game there are numbers already preset on the board which can not be changed. From there, the goal is to fill the empty spaces with the right numbers. Each line going horizontally and vertically, as well as the 3x3  subset blocks, must have all numbers 1 - 9 without any repeats. 

## Overview:
 In an attempt to build off the principles we learned this semester, we proposed the idea of creating a speedy sudoku solver that will use heuristics and backtracking to solve a sudoku puzzle as fast as possible. By predicting out which squares are most likely to be solvable, less time will be wasted using the typical depth-first search algorithm, which is similar to brute forcing the problem in this usage. By using backtracking, we were able to make a decision tree of sorts, trying all possible options for a square and recursively calling backtrack on it until only one branch is possible. By using the heuristics in tandem to backtracking, we figured out the selecting to backtrack on squares with the fewest remaining options left, causing failure faster and allowing the algorithm to proceed. 

## Experiment Setup:
We created our Speedy Sudoku Solver to be run within the Linux command line. We initially were using Python as our choice of language in the various implementations, but due to the size of the problem space, and because we were originally considering the idea of running various parts of this project in parallel with each other, we decided to write our AI in Golang. This was also a good choice because this language is a good deal faster than Python. Because we needed to very accurately measure the performance of a particular implementation, we decided to utilize the built-in Golang ‘time’ library as our choice of performance benchmarking tool. Each test was run 10 times on10 different puzzles (100 total), and the times presented in this paper are an average of these runs. Additionally, we created 5 levels of tests (10 of each) to measure how performance changed across different difficulties. We stuck to only 9x9 sudoku puzzles that ranged from ‘easy’ to ‘evil’ difficulties.

## Naive Implementation:
Our very first implementation took a naive approach to solving sudoku. An inexperienced programmer may have looked at the problem of solving a sudoku puzzle and thought that there should always be a way to solve a puzzle, given the three main rules:

One of each number value (1-9) in each row.
One of each number value (1-9) in each column.
One of each number value (1-9) in each 3x3 box.

By following these rules, we can easily solve puzzles of  ‘easy’ to ‘medium’ level difficulty, but in order to solve very hard problems, there are often much higher levels of abstraction that an AI would have to make in order to solve them. There are no more rules at a harder level, but there are less values placed to help guarantee that there will be a square with only 1 possible value for it at a given game state.

## Backtracking:
In order to solve harder puzzles, we have to either build on more sophisticated ways of detecting if there is a value that can be placed, or we have to change our tactic completely. We chose the latter. We now introduce the topic of backtracking. Consider a decision tree, similar to the ones used in Chess AI’s. Typically the AI will make a tree of all possible moves, and all possible responses, and all possible moves to counter those responses, and so on. The AI will then look, say 5 moves ahead and pursue the best games state available based on the current state of the game. We can apply a similar decision tree method with Sudoku, except now we do not have to take into account the unpredictability of an opponent, and there is only one final state to which we want to arrive at: the solution. 

Using the same depth-first search algorithm as in the original naive version, we can arrive at the first square that needs to be solved, and recur on all of those values. For example, say the first empty square has 3 possible values. We can recur on each possible value and continue playing until we reach a state where there is no possible value in a square. Its called back tracking because we explore all possible branches, and all of their sub branches until we reach a winning or losing game state. If it’s a losing one, we backtrack to where it was last successful and explore all of the remaining options of values.

## Priority Queue With Heuristics:
 Backtracking provided us with a reliable solution, but as you could probably see, if you use the depth-first approach as we have up until now, the amount of branches you need to search becomes massive. Our goal of this AI is to make the best choices of what to solve next, and so we now implement heuristics. The first heuristic we decided to use was how many possible values can be placed in a given square. This notion is fairly simple, if there are less values, then you have to explore less branches. More than this, though, is that if you start with the squares that have the least possible values, you can quickly start placing the easy ones on the board, making all of the other squares have less and less values. If you could reliably solve the entire puzzles where you only explore squares with 1, 2, or a seldom 3 branches to explore, then the puzzle suddenly becomes easy again. But how will we assess the board without wasting more time than we are saving with it?

Our original idea would be to re-assess the board every time we recur, to get the most up-to-date optimization. We implemented this by creating a second 2D array (a second board) and instead of values and empty squares, we now had the number of possible values that could be placed in the corresponding location in the real board. This was a very accurate solution in terms of prioritizing, but it ended up increasing the time it took to solve a puzzle by a factor of 5. So we had to re-think our approach. We decided to use heuristics to assess the board at the very beginning and only at the beginning in order to reliably predict the best way to solve the board with only one assessment. We did this by creating two arrays, one that held the number of possible values, and another that held encoded values that represented locations on the board. We would now return the list of locations and would not need to search for the next best every time because we had a priority queue that would tell exactly what to solve next.

## Choosing Which Values to Try First:
By reordering the numbers we try at the beginning, then we can backtrack less by choosing to try the values that are closer to being solvable first. We will base this weight on how many of this given number has been places so far.

## Putting It All Together:
We will now show you the results of our major improvements. The naive implementation is the original depth-first search attempt of going through the board and only solving a square when there is only one possible answer. The backtracking solution still uses depth-first search but now ensures that every square gets solved. The priority queue implementation takes all previous improvements and searches for squares that are more likely to be solved and tries those first. This prioritization work is done ahead of time. The optimal found solution is a combination of all previous solutions with the heuristics we found to be helpful. 

## Results:
Naive Backtracking - Evil - 18.430 ms (0.02 seconds)
Priority Queue (Based on Simple Heuristics) - Evil - 2.684 ms (0.002 seconds)
Optimal Solution (With Value Ordering) - Evil - 146.327 μs (0.0001 seconds)

## If We had More Time:
If we had more time to increase the speed or capability of our AI, we would possible consider some of the following ideas:
Bigger puzzles.
Reading the image of a board to avoid having to manually input the puzzle.
Storing the board as a dictionary for constant time lookup as opposed to N^2 space complexity.
More heuristics on bigger boards.
Using machine learning on many boards to identify mathematical or geographical advantages for given board formations or number placements.
Showing the work done by the AI in a GUI that would visually represent what the program did/tried to solve the puzzle.
Unfortunately, due to lack of time, additional responsibilities, and guidance to only use 3 projects worth of work (no more, no less) we were unable to get to all of these interesting ideas for further improvement. 
