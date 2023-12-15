package main

import (
	"fmt"
	"math"
	"slices"
)

type Cell struct {
	Row            int
	Column         int
	Box            int
	PossibleValues []byte
}

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

func solveSudoku(board [][]byte) {
	findNumber(board)
	// for !isSolved(board) {
	// }
}

// func isSolved(board [][]byte) bool {
// 	for i := 0; i < len(board); i++ {
// 		for j := 0; j < len(board[i]); j++ {
// 			if board[i][j] == '.' {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

func findNumber(board [][]byte) {
	var cells []Cell
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				existingValues := []byte{}

				rowNumbers := getRowNumbers(board, i, existingValues)
				columnNumbers := getColumnNumbers(board, j, existingValues)
				boxNumbers, boxNum := getBoxNumbers(board, i, j, existingValues)

				cells = append(cells, Cell{Row: i, Column: j, Box: boxNum, PossibleValues: possibleValues})

				fmt.Println(possibleValues)
				return

				// if len(numbers) == 8 {
				// 	var sum int
				// 	for k, _ := range numbers {
				// 		sum += int(k)
				// 	}
				// 	board[i][j] = byte(45 - sum + 48)
				// 	fmt.Println()
				// 	return
				// }
			}
		}
	}

}

func getRowNumbers(board [][]byte, rowIndex int, numbers []byte) []byte {
	var values []byte
	for _, b := range board[rowIndex] {
		if b != '.' && !slices.Contains(numbers, b) {
			values = append(values, b)
		}
	}
}

func getBoxNumbers(board [][]byte, row int, column int, numbers []byte) ([]byte, int) {
	var boxRow int
	var boxColumn int
	switch {
	case row < 3:
		boxRow = 1
	case row < 6:
		boxRow = 2
	default:
		boxRow = 3
	}

	switch {
	case column < 3:
		boxColumn = 0
	case row < 6:
		boxColumn = 1
	default:
		boxColumn = 2
	}

	boxNumber := boxColumn*3 + boxRow

	boxY := int(math.Ceil(float64(row+1)/3)*3 - 3)
	boxX := int(math.Ceil(float64(column+1)/3)*3 - 3)
	var values []byte

	for i := boxY; i <= boxY+2; i++ {
		for j := boxX; j <= boxX+2; j++ {
			curr := board[i][j]
			if curr != '.' && !slices.Contains(numbers, curr) {
				values = append(values, curr)
			}
		}
	}
	return values, boxNumber
}

func getColumnNumbers(board [][]byte, columnIndex int, numbers []byte) []byte {
	var values []byte
	for _, row := range board {
		val := row[columnIndex]
		if val != '.' && !slices.Contains(numbers, b) {
			values = append(values, val)
		}
	}
}

// func getColumns(board [][]byte) [][]byte {
// 	var columns [][]byte
// 	for j := 0; j < len(board[0]); j++ {
// 		curr := []byte{}
// 		for i := range board {
// 			curr = append(curr, board[i][j])
// 		}
// 		columns = append(columns, curr)
// 	}

// 	return columns
// }

// func getBoxes(board [][]byte) [][]byte {
// 	// boxY := int(math.Ceil(float64(row+1)/3)*3 - 3)
// 	// boxX := int(math.Ceil(float64(column+1)/3)*3 - 3)

// 	for i := 0; i <= 2; i++ {
// 		for j := 0; j <= 2; j++ {
// 			curr := board[i][j]
// 			if curr != '.' {
// 				numbers[curr-48] = true
// 			}
// 		}
// 	}
// }
