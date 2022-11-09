package main

import (
	"os"

	"github.com/01-edu/z01"
)

func intervalThree(x int) []int {
	if x == 0 || x == 1 || x == 2 {
		return []int{0, 1, 2}
	}
	if x == 3 || x == 4 || x == 5 {
		return []int{3, 4, 5}
	}
	if x == 6 || x == 7 || x == 8 {
		return []int{6, 7, 8}
	}
	return nil
}
func isNumberAllowed(board [][]rune, coords [2]int, val rune) bool {
	for x, v := range board[coords[1]] {
		if x != coords[0] && v == val {
			return false
		}
	}
	for y := 0; y < 9; y++ {
		if y != coords[1] && board[y][coords[0]] == val {
			return false
		}
	}
	for _, y := range intervalThree(coords[1]) {
		for _, x := range intervalThree(coords[0]) {
			if board[y][x] == val && !(y == coords[1] && x == coords[0]) {
				return false
			}
		}
	}
	return true
}
func isBoardValid(board [][]rune) bool {
	if len(board) != 9 {
		return false
	}
	for y, row := range board {
		if len(row) != 9 {
			return false
		}
		for x, e := range row {
			if (e < '1' || e > '9') && e != '.' {
				return false
			}
			if e != '.' && !isNumberAllowed(board, [2]int{x, y}, e) {
				return false
			}
		}
	}
	return true
}
func getEmptyBoxes(board [][]rune) [][]int {
	var boxes [][]int
	for y, row := range board {
		for x, r := range row {
			if r == '.' {
				boxes = append(boxes, []int{x, y})
			}
		}
	}
	return boxes
}
func solve(board [][]rune, emptyBoxes [][]int, indexBoxToFill int) bool {
	if indexBoxToFill >= len(emptyBoxes) {
		return true
	}
	x := emptyBoxes[indexBoxToFill][0]
	y := emptyBoxes[indexBoxToFill][1]
	for nb := '1'; nb <= '9'; nb++ {
		if isNumberAllowed(board, [2]int{x, y}, nb) {
			board[y][x] = nb
			if solve(board, emptyBoxes, indexBoxToFill+1) {
				return true
			}
			board[y][x] = '.'
		}
	}
	return false
}
func printBoard(board [][]rune) {
	for _, row := range board {
		for i, r := range row {
			if i != 0 {
				z01.PrintRune(' ')
			}
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
func main() {
	var board [][]rune
	for _, row := range os.Args[1:] {
		board = append(board, []rune(row))
	}
	if !isBoardValid(board) {
		println("Error")
		return
	}
	emptyBoxes := getEmptyBoxes(board)
	if !solve(board, emptyBoxes, 0) {
		println("Error")
	}
	printBoard(board)
}
