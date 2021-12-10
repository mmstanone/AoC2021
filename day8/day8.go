package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func task_1() {
	file, _ := os.Open("puzzle8.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maps := map[int]int{2: 1, 4: 4, 3: 7, 7: 8}
	counts := make(map[int]int)
	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), "|")
		for _, elem := range strings.Split(strings.TrimSpace(numbers[1]), " ") {
			if mapping, ok := maps[len(strings.TrimSpace(elem))]; ok {
				counts[mapping] += 1
			}

		}
	}
	fmt.Println(sum_map(counts))
}

func sum_map(m map[int]int) int {
	sum := 0
	for _, elem := range m {
		sum += elem
	}
	return sum
}

func filter_length(slice []string, length int) []string {
	filtered := []string{}

	for i := range slice {
		if len(slice[i]) == length {
			filtered = append(filtered, slice[i])
		}
	}
	return filtered
}

func exclude_from_string(slice []string, exclude string) []string {
	filtered := []string{}

	for i := range slice {
		if slice[i] == exclude {
			continue
		}
		filtered = append(filtered, slice[i])

	}
	return filtered
}

func find_a(mappings *map[string]string, number_mappings map[int]map[string]struct{}) {
	for a := range number_mappings[7] {
		if _, ok := number_mappings[1][a]; !ok {
			(*mappings)["a"] += a
			(*mappings)["left"] = strings.Replace((*mappings)["left"], a, "", 1)
		}
	}
}

func map_contains(number_mappings map[string]struct{}, a string) bool {
	for elem := range number_mappings {
		if elem == a {
			return true
		}
	}
	return false
}

func poop(candidates []string, mappings map[string]string, number_mappings *map[int]map[string]struct{}, number int, char string) []string {
	var final []string
	for _, candidate := range candidates {
		if strings.Contains(candidate, mappings[char]) {
			imm_res := make(map[string]struct{})
			for _, elem := range candidate {
				imm_res[string(elem)] = struct{}{}
			}
			(*number_mappings)[number] = imm_res
			final = exclude_from_string(candidates, candidate)
			return final
		}
	}
	return final
}

func find_eg(mappings *map[string]string, maps []string, number_mappings *map[int]map[string]struct{}) {
	for a := range (*number_mappings)[8] {
		if _, ok := (*number_mappings)[4][a]; !ok {
			if a == (*mappings)["a"] {
				continue
			}
			(*mappings)["e"] += a
			(*mappings)["g"] += a
			(*mappings)["left"] = strings.Replace((*mappings)["left"], a, "", 1)
		}
	}
	candidates_235 := filter_length(maps, 5)

	e_candidates, _ := (*mappings)["e"]

	for _, char := range e_candidates {
		for _, candidate := range candidates_235 {
			if !strings.Contains(candidate, string(char)) {
				(*mappings)["e"] = string(char)
				(*mappings)["g"] = strings.Trim(e_candidates, string(char))
			}
		}
	}
	all_present := make(map[string]int)
	for _, candidate := range candidates_235 {
		for _, char := range candidate {
			if string(char) != (*mappings)["e"] {
				all_present[string(char)] += 1
			}
		}
	}

	for a, val := range all_present {
		if val == 1 {
			(*mappings)["b"] += a
			(*mappings)["left"] = strings.Replace((*mappings)["left"], a, "", 1)
		}
	}
	candidates_23 := poop(candidates_235, *mappings, number_mappings, 5, "b")

	all_present = make(map[string]int)
	for _, candidate := range candidates_23 {
		for _, char := range candidate {
			if string(char) != (*mappings)["e"] {
				all_present[string(char)] += 1
			}
		}
	}

	for a, val := range all_present {
		if val == 1 {
			(*mappings)["f"] += a
			(*mappings)["left"] = strings.Replace((*mappings)["left"], a, "", 1)
		}
	}

	candidates_2 := poop(candidates_23, *mappings, number_mappings, 2, "a")

	for a := range (*number_mappings)[2] {
		if _, ok := (*number_mappings)[5][a]; !ok {
			if a != (*mappings)["e"] {
				(*mappings)["c"] = a
				(*mappings)["left"] = strings.Replace((*mappings)["left"], a, "", 1)
			}
		}

	}
	poop(candidates_2, *mappings, number_mappings, 3, "f")
	(*mappings)["d"] = (*mappings)["left"]
	(*mappings)["left"] = ""
}

func get_mapped(char string, mappings map[string]string) string {
	for k, elem := range mappings {
		if elem == char {
			return k
		}
	}
	return char
}

func find_real(real_mappings map[int]map[string]struct{}, maps map[string]struct{}) int {
	for k, v := range real_mappings {
		if len(maps) != len(v) {
			continue
		}
		u := 1
		for x := range maps {
			if _, ok := v[x]; !ok {
				u = 0
				break
			}
		}
		if u == 1 {
			return k
		}
	}
	return -1
}

func task_2() {
	file, _ := os.Open("puzzle8.txt")
	defer file.Close()

	basic_mappings := map[int]int{2: 1, 4: 4, 3: 7, 7: 8}

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {

		/*
		    aaaa
		   b    c
		   b    c
		    dddd
		   e    f
		   e    f
		    gggg
		*/
		mappings := map[string]string{"a": "", "b": "", "c": "", "d": "", "e": "", "f": "", "g": "", "left": "abcdefg"}
		real_mappings := map[int]map[string]struct{}{0: {"a": {}, "b": {}, "c": {}, "e": {}, "f": {}, "g": {}},
			1: {"c": {}, "f": {}},
			2: {"a": {}, "c": {}, "d": {}, "e": {}, "g": {}},
			3: {"a": {}, "c": {}, "d": {}, "f": {}, "g": {}},
			4: {"b": {}, "c": {}, "d": {}, "f": {}},
			5: {"a": {}, "b": {}, "d": {}, "f": {}, "g": {}},
			6: {"a": {}, "b": {}, "d": {}, "e": {}, "f": {}, "g": {}},
			7: {"a": {}, "c": {}, "f": {}},
			8: {"a": {}, "b": {}, "c": {}, "d": {}, "e": {}, "f": {}, "g": {}},
			9: {"a": {}, "b": {}, "c": {}, "d": {}, "f": {}, "g": {}}}
		number_mappings := map[int]map[string]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}, 0: {}}
		numbers := strings.Split(scanner.Text(), "|")
		maps := strings.Split(strings.TrimSpace(numbers[0]), " ")
		nums := strings.TrimSpace(numbers[1])
		for _, elem := range maps {
			if mapping, ok := basic_mappings[len(strings.TrimSpace(elem))]; ok {
				for _, char := range elem {
					number_mappings[mapping][string(char)] = struct{}{}
				}
			}
		}

		find_a(&mappings, number_mappings)
		find_eg(&mappings, maps, &number_mappings)
		num := 0
		for _, number := range strings.Split(nums, " ") {
			used := make(map[string]struct{})
			for _, char := range number {
				used[get_mapped(string(char), mappings)] = struct{}{}
			}
			n := find_real(real_mappings, used)
			num = num*10 + n
		}
		sum += num
		// return
	}
	fmt.Println(sum)
}

func main() {
	// task_1()
	task_2()

}
