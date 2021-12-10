package utils

import (
	"os"
	"strconv"
	"strings"
)

func ParseCsv(file string) []string {
	f, _ := os.ReadFile(file)
	return strings.Split(strings.Trim(string(f), "\n "), ",")
}

func ConvertToInts(vals []string) []int {
	res := make([]int, 0)
	for _, val := range vals {
		conv, _ := strconv.Atoi(val)
		res = append(res, conv)
	}
	return res
}
