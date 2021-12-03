package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func leq(x, y float64) bool {
	return x <= y
}

func geq(x, y float64) bool {
	return x >= y
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func print_results(a string, b string) {
	gamma, err := strconv.ParseInt(a, 2, 32)
	check(err)
	eps, err := strconv.ParseInt(b, 2, 32)
	check(err)
	fmt.Println(gamma, eps, gamma*eps)

}

func get_cmp(arr []string, index int, comparator func(float64, float64) bool, ok byte, nok byte) byte {
	curr := 0
	for ind := 0; ind < len(arr); ind++ {
		if arr[ind][index] == ok {
			curr++
		}
	}

	if comparator(float64(curr), float64(len(arr))/float64(2)) {
		return ok
	} else {
		return nok
	}
}

func cross_out_rows(slice *[]string, index int, compared byte) {
	i := 0
	for len(*slice) > 1 && i < len(*slice) {
		if (*slice)[i][index] != compared {
			*slice = remove(*slice, i)
		} else {
			i += 1
		}
	}

}

func main() {
	file, e := os.Open("puzzle3.txt")
	check(e)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fins []string
	var negs []string
	for scanner.Scan() {
		current := scanner.Text()
		fins = append(fins, current)
		negs = append(negs, current)
	}

	for index := 0; index < len(fins[0]); index++ {
		fins_cmp := get_cmp(fins, index, geq, '1', '0')
		neg_cmp := get_cmp(negs, index, leq, '0', '1')

		cross_out_rows(&fins, index, fins_cmp)
		cross_out_rows(&negs, index, neg_cmp)
		fmt.Println(len(fins), len(negs), fins_cmp, neg_cmp)

	}

	if len(fins) > 1 || len(negs) > 1 {
		panic("wtf fucked up")
	}

	print_results(fins[0], negs[0])
}
