package main

import (
	"bufio"
	"fmt"
	// "io"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func task01() {
	f, err := os.Open("puzzle1.txt")
	check(err)
	defer f.Close()

	inc_counter := 0
	scanner := bufio.NewScanner(f)
	last_scan := -1
	for scanner.Scan() {
		current_scan, err := strconv.Atoi(scanner.Text())
		check(err)
		if last_scan == -1 {
			last_scan = current_scan
			continue
		}

		if last_scan < current_scan {
			inc_counter++
		}

		last_scan = current_scan
	}

	fmt.Printf("Increasing: %d\n", inc_counter)

}
func main() {
	buff := make([]int, 0)

	f, err := os.Open("puzzle1.txt")
	check(err)
	defer f.Close()

	inc_counter := 0
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		current_scan, err := strconv.Atoi(scanner.Text())
		check(err)
		buff = append(buff, current_scan)
	}

	last_sum := -1
	for i := 0; i < len(buff)-2; i++ {
		curr_sum := 0
		for j := 0; j < 3; j++ {
			curr_sum += buff[i+j]
		}
		if last_sum == -1 {
			last_sum = curr_sum
			continue
		}
		if last_sum < curr_sum {
			inc_counter += 1
		}
		last_sum = curr_sum
	}
	fmt.Printf("Increasing: %d\n", inc_counter)

}
