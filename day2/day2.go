package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func task01(scanner bufio.Scanner) {
	depth := 0
	location := 0

	for scanner.Scan() {
		line := scanner.Text()
		sp := strings.Fields(line)

		direction := sp[0]
		movement, err := strconv.Atoi(sp[1])
		check(err)

		switch direction {
		case "forward":
			location += movement
		case "up":
			depth -= movement
		case "down":
			depth += movement
		default:
			panic("lol")
		}
	}
	fmt.Println(depth * location)

}

func main() {
	fp, err := os.Open("puzzle2.txt")
	check(err)
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	depth := 0
	location := 0
	aim := 0

	for scanner.Scan() {
		line := scanner.Text()
		sp := strings.Fields(line)

		direction := sp[0]
		movement, err := strconv.Atoi(sp[1])
		check(err)

		switch direction {
		case "forward":
			location += movement
			depth += movement * aim
		case "up":
			aim -= movement
		case "down":
			aim += movement
		default:
			panic("lol")
		}
	}
	fmt.Println(depth * location)
}
