package main

import "fmt"

// medium

func isValidSudoku(board [][]byte) bool {
	var rows, cols, boxes [9][9]bool // Это гораааздо лучше чем слайсы как ниже

	//rows := make([][]bool, 9)
	//cols := make([][]bool, 9)
	//boxes := make([][]bool, 9)
	//
	//for i := 0; i < 9; i++ {
	//	rows[i] = make([]bool, 9)
	//	cols[i] = make([]bool, 9)
	//	boxes[i] = make([]bool, 9)
	//}


	for i := 0; i < 9; i++ { // i строка
		for j := 0; j < 9; j++ { // j-тый столбец
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - '1'

			if rows[i][num] || cols[j][num] || boxes[(i/3)*3+(j/3)][num] { // and boxes
				return false
			}

			rows[i][num] = true
			cols[j][num] = true
			boxes[(i/3)*3+(j/3)][num] = true
		}
	}

	return true
}

func main() {
	var a byte
	a = '3'

	fmt.Println(a)
}
