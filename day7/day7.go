package main

import (
	// this gives me nothing anyway T_T
	"aoc/utils"
	"fmt"
	"math"
)

func main() {
	parsed := utils.ParseCsv("puzzle7.txt")
	vals := utils.ConvertToInts(parsed)
	mean := float64(0)
	for _, val := range vals {
		mean += float64(val)
	}

	mean /= float64(len(vals))
	fmt.Println(mean)

	stddev := float64(0)

	for _, val := range vals {
		stddev += (float64(val) - mean) * (float64(val) - mean)
	}

	variance := stddev / float64(len(vals))
	stddev = math.Sqrt(variance)

	min_sum := math.MaxInt32
	// haha go go fast
	for meeting_point := 0; meeting_point < len(vals); meeting_point++ {
		sum := 0
		for _, val := range vals {
			for i := 1; i < int(math.Abs(float64(val-meeting_point)))+1; i++ {
				sum += i
			}
		}

		if sum < min_sum {
			min_sum = sum
		}
	}
	fmt.Println(min_sum)
}
