package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func leq(a int, b int) bool {
	return a <= b
}

func geq(a int, b int) bool {
	return a >= b
}

func print_field(field [][]int) {
	for _, row := range field {
		for _, elem := range row {
			if elem == 0 {
				fmt.Printf(". ")
			} else {
				fmt.Printf("%d ", elem)
			}
		}
		fmt.Printf("\n")
	}
}

func minmax(a int, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func choose_steps(a int, b int) (int, int, int, func(int, int) bool) {
	var move int
	var start int
	var end int
	var comp func(int, int) bool
	min, max := minmax(a, b)
	if a < b {
		move = 1
		start = min
		end = max
		comp = leq
	} else if a > b {
		move = -1
		start = max
		end = min
		comp = geq
	} else {
		move = 0
		start = min
		end = min
		comp = geq
	}
	return move, start, end, comp
}

func main() {
	file, _ := os.Open("puzzle5.txt")

	defer file.Close()
	const SIZE = 1000

	field := make([][]int, SIZE)

	for ind, _ := range field {
		field[ind] = make([]int, SIZE)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "->")

		var first_coords []int
		var second_coords []int
		for _, val := range strings.Split(line[0], ",") {
			v, _ := strconv.Atoi(strings.Trim(val, " "))
			first_coords = append(first_coords, v)
		}
		for _, val := range strings.Split(line[1], ",") {
			v, _ := strconv.Atoi(strings.Trim(val, " "))
			second_coords = append(second_coords, v)
		}

		c1, r1 := first_coords[0], first_coords[1]
		c2, r2 := second_coords[0], second_coords[1]

		c_move, c_start, c_end, comp_col := choose_steps(c1, c2)
		r_move, r_start, r_end, comp_row := choose_steps(r1, r2)

		c := c_start
		r := r_start
		for comp_col(c, c_end) && comp_row(r, r_end) {
			field[r][c] += 1
			c += c_move
			r += r_move
		}
	}

	score := 0
	for _, row := range field {
		for _, col := range row {
			if col >= 2 {
				score += 1
			}
		}
	}
	fmt.Println(score)
}
