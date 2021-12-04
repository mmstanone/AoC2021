package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bingo interface {
	cross_out(number int)
	did_win()
	score(number int)
}

type BingoField struct {
	value   int
	crossed bool
}

type BingoBoard struct {
	board [][]BingoField
	won   bool
}

func (b *BingoBoard) cross_out(number int) (int, int) {
	for row_ind, row := range (*b).board {
		for col_ind, field := range row {
			if field.value == number {
				(*b).board[row_ind][col_ind].crossed = true
				return row_ind, col_ind
			}
		}
	}
	return -1, -1
}

func (b *BingoBoard) did_win() bool {
	cols := make([]int, 5)
	for _, row := range b.board {
		count := 0
		for col, field := range row {
			if field.crossed {
				count++
				cols[col]++
			}
		}
		if count == 5 {
			return true
		}
	}
	for _, elem := range cols {
		if elem == 5 {
			return true
		}
	}
	return false
}

func (b *BingoBoard) score(number int) int {
	sc := 0
	for _, row := range b.board {
		for _, field := range row {
			if !field.crossed {
				sc += field.value
			}
		}
	}
	return sc * number
}

func main() {
	file, _ := os.Open("puzzle4.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []int

	scanner.Scan()

	line := strings.Split(scanner.Text(), ",")

	for _, char := range line {
		conv, _ := strconv.Atoi(char)
		numbers = append(numbers, conv)
	}

	fmt.Println(numbers)

	var boards []BingoBoard

	curr_board := make([][]BingoField, 0)

	for scanner.Scan() {
		if len(scanner.Text()) < 3 {
			boards = append(boards, BingoBoard{curr_board, false})
			curr_board = make([][]BingoField, 0)
			continue
		}

		curr_line := strings.Split(scanner.Text(), " ")
		curr_row := make([]BingoField, 0)
		for _, char := range curr_line {
			conv, _ := strconv.Atoi(char)
			if conv == 0 {
				continue
			}
			curr_row = append(curr_row, BingoField{conv, false})
		}
		curr_board = append(curr_board, curr_row)
	}
	boards = append(boards, BingoBoard{curr_board, false})

	boards = boards[1:]

	for _, number := range numbers {
		for ind, board := range boards {
			if board.won {
				continue
			}
			board.cross_out(number)
			if boards[ind].did_win() {
				boards[ind].won = true
				count := 0
				for _, _board := range boards {
					if !_board.won {
						count++
					}
				}
				if count == 0 {
					fmt.Println(number)
					fmt.Println(board.score(number))
					return
				}
			}
		}
	}
}
