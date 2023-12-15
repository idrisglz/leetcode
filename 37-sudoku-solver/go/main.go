package main

func main() {
	var board = [][]byte{
		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
	}

	solveSudoku(board)
}

func isEligible(board [][]byte, row, col int, num byte) bool {
	for x := 0; x < 9; x++ {
		if board[row][x] == num {
			return false
		}
		if board[x][col] == num {
			return false
		}
		if board[3*(row/3)+x/3][3*(col/3)+x%3] == num {
			return false
		}
	}
	return true
}

func solveSudoku(board [][]byte) bool {
	row, col := -1, -1
	isComplete := false
outer_loop:
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				row, col = i, j
				isComplete = true
				break outer_loop
			}
		}
	}
	if !isComplete {
		return true
	}
	for num := byte('1'); num <= byte('9'); num++ {
		if isEligible(board, row, col, num) {
			board[row][col] = num
			if solveSudoku(board) {
				return true
			}
			board[row][col] = '.'
		}
	}
	return false
}
