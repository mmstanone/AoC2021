package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("puzzle6.txt")

	values := strings.Split(string(input), ",")
	lives := make([]int, 9)
	lost_lives := make([]int, 9)
	for _, val := range values {
		conv, _ := strconv.Atoi(strings.Trim(val, "\n"))
		lives[conv]++
	}

	fmt.Println(lives)
	for iterations := 0; iterations < 256; iterations++ {
		copy(lost_lives, lives)
		lives = make([]int, 9)
		for ind := 0; ind < 9; ind++ {
			if ind == 0 {
				lives[8] = lost_lives[0]
				lives[6] = lost_lives[0]
			} else {
				lives[ind-1] += lost_lives[ind]
			}

		}
		count := 0
		for ind := 0; ind < 9; ind++ {
			count += lives[ind]
		}
		fmt.Println(count)
	}

}
