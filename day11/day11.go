package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type heatmap_point struct {
	value   int
	visited bool
}

func increase_value(heatmap *[][]heatmap_point, i int, j int) {
	if !(0 <= j && j < len((*heatmap)) && 0 <= i && i < len((*heatmap)[j])) {
		return
	}
	(*heatmap)[i][j].value += 1
}

func step(heatmap *[][]heatmap_point) int {
	for i := range *heatmap {
		for j := range (*heatmap)[i] {
			increase_value(heatmap, i, j)
		}
	}

	for i := range *heatmap {
		for j := range (*heatmap)[i] {
			dfs(heatmap, j, i)
		}
	}
	count := 0
	for i := range *heatmap {
		for j := range (*heatmap)[i] {
			if (*heatmap)[i][j].value > 9 {
				(*heatmap)[i][j].value = 0
				count += 1
			}
			(*heatmap)[i][j].visited = false
		}
	}
	return count
}

func do_something_with(heatmap *[][]heatmap_point, y int, x int, something func(heatmap *[][]heatmap_point, y int, x int)) {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			something(heatmap, y+dy, x+dx)
		}
	}
}

func dfs(heatmap *[][]heatmap_point, y int, x int) {
	if !(0 <= y && y < len((*heatmap)) && 0 <= x && x < len((*heatmap)[y])) {
		return
	}
	if (*heatmap)[y][x].value <= 9 || (*heatmap)[y][x].visited {
		return
	}
	(*heatmap)[y][x].visited = true

	do_something_with(heatmap, y, x, increase_value)
	do_something_with(heatmap, y, x, dfs)
	return
}

func main() {
	file, _ := os.Open("puzzle11.txt")
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

	i := 1
	count := 0
	for i < 300 {
		res := step(&heatmap)
		if res == 100 {
			fmt.Println(i)
		}
		count += res
		i += 1
	}
	println(count)

}
