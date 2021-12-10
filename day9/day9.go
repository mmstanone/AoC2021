package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type d struct {
	x int
	y int
}

func is_lowest(heatmap [][]heatmap_point, x int, y int) bool {
	ways := []d{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	count := 0
	all_count := 0
	for _, way := range ways {
		if 0 <= way.x+x && way.x+x < len(heatmap[y]) && 0 <= way.y+y && way.y+y < len(heatmap) {
			all_count++
			if heatmap[way.y+y][way.x+x].value > heatmap[y][x].value {
				count++
			}
		}
	}

	return count == all_count
}

type point struct {
	x int
	y int
}

type heatmap_point struct {
	value   int
	visited bool
}

func dfs(heatmap *[][]heatmap_point, x int, y int, previous_val int) int {
	if !(0 <= y && y < len((*heatmap)) && 0 <= x && x < len((*heatmap)[y])) {
		return 0
	}

	if (*heatmap)[y][x].value < previous_val {
		return 0
	}

	if (*heatmap)[y][x].visited || (*heatmap)[y][x].value == 9 {
		return 0
	}

	(*heatmap)[y][x].visited = true

	size := 1
	ways := []d{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for _, way := range ways {
		size += dfs(heatmap, x+way.x, y+way.y, (*heatmap)[y][x].value)
	}
	return size
}

func main() {
	file, _ := os.Open("puzzle9.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var heatmap [][]heatmap_point

	for scanner.Scan() {
		current := strings.Split(strings.TrimSpace(scanner.Text()), "")
		var n []heatmap_point
		for _, elem := range current {
			conv, _ := strconv.Atoi(elem)
			n = append(n, heatmap_point{conv, false})
		}
		heatmap = append(heatmap, n)
	}

	count := 0
	var points []point
	for y := range heatmap {
		for x := range heatmap[y] {
			if is_lowest(heatmap, x, y) {
				points = append(points, point{x, y})
				count += 1 + heatmap[y][x].value
			}
		}
	}
	fmt.Println(count)
	fmt.Println(points)
	sums := make([]int, 0)
	for _, point := range points {
		sums = append(sums, dfs(&heatmap, point.x, point.y, heatmap[point.y][point.x].value))
		for y := range heatmap {
			for x := range heatmap[y] {
				heatmap[y][x].visited = false
			}
		}
	}

	sort.Slice(sums, func(i, j int) bool {
		return sums[j] < sums[i]
	})

	out := 1
	for _, elem := range sums[:3] {
		out *= elem
	}

	fmt.Println(out)
}
