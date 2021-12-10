package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func parse(line string) int {
	closures := map[byte]byte{'>': '<', ')': '(', '}': '{', ']': '['}
	stack := make([]byte, 0)
	values := map[byte]int{')': 3, ']': 57, '}': 1197, '>': 25137}

	for _, char := range line {
		if closure, ok := closures[byte(char)]; ok {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if last != closure {
				return values[byte(char)]
			}
		} else {
			stack = append(stack, byte(char))
		}
	}
	return 0
}

func reverse(slice []byte) []byte {
	new := make([]byte, 0)

	for i := len(slice) - 1; i >= 0; i-- {
		new = append(new, slice[i])
	}

	return new
}

func parse2(line string) int {
	closures := map[byte]byte{'>': '<', ')': '(', '}': '{', ']': '['}
	stack := make([]byte, 0)
	values := map[byte]int{'(': 1, '[': 2, '{': 3, '<': 4}

	res := 0

	for _, char := range line {
		if _, ok := closures[byte(char)]; ok {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, byte(char))
		}
	}

	reversed := reverse(stack)
	for _, char := range reversed {
		res *= 5
		res += values[byte(char)]
	}
	return res
}

func main() {
	file, _ := os.Open("puzzle10.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scores := make([]int, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		res := parse(line)
		if res != 0 {
			continue
		}

		scores = append(scores, parse2(line))
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})
	fmt.Println(sort.IntSlice(scores)[(len(scores)-1)/2])
}
