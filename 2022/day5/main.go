package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func single_shift(start_stack, end_stack []string) ([]string, []string) {

	end_stack = append(end_stack, start_stack[len(start_stack)-1])
	n := len(start_stack) - 1
	start_stack = start_stack[:n]

	return start_stack, end_stack
}

func multi_shift(start_stack []string, end_stack []string, move_count int) ([]string, []string) {

	start_crate := len(start_stack) - move_count
	top_crate := start_crate + move_count
	crate_stack := start_stack[start_crate:top_crate]
	end_stack = append(end_stack, crate_stack...)

	n := len(start_stack) - move_count

	start_stack = start_stack[:n]

	return start_stack, end_stack
}

func cratemover_9000(containers_map map[int][]string, instructions []string) map[int][]string {
	for _, steps := range instructions {
		step := strings.Split(steps, " ")
		move_count, _ := strconv.Atoi(step[1])
		start_stack, _ := strconv.Atoi(step[3])
		end_stack, _ := strconv.Atoi(step[5])

		for i := 0; i < move_count; i++ {
			containers_map[start_stack], containers_map[end_stack] = single_shift(containers_map[start_stack], containers_map[end_stack])
		}

	}
	return containers_map
}

func cratemover_9001(containers_map map[int][]string, instructions []string) map[int][]string {
	for _, steps := range instructions {
		step := strings.Split(steps, " ")
		move_count, _ := strconv.Atoi(step[1])
		start_stack, _ := strconv.Atoi(step[3])
		end_stack, _ := strconv.Atoi(step[5])

		if move_count == 1 {
			containers_map[start_stack], containers_map[end_stack] = single_shift(containers_map[start_stack], containers_map[end_stack])
		} else {
			containers_map[start_stack], containers_map[end_stack] = multi_shift(containers_map[start_stack], containers_map[end_stack], move_count)
		}

	}
	return containers_map
}
func main() {
	file, err := os.ReadFile("../inputs/day5.txt")
	if err != nil {
		fmt.Println(err)
	}

	containers_set := strings.Split(string(file), "\n\n")[0]
	instruction_set := strings.Split(string(file), "\n\n")[1]

	var containers_map = make(map[int][]string)
	container_rows := strings.Split(string(containers_set), "\n")
	stack_count_row := strings.Replace(container_rows[len(container_rows)-1], " ", "", -1)
	stack_count := len(strings.Split(stack_count_row, ""))

	for i := 1; i <= stack_count; i++ {
		var containers []string
		containers_map[i] = containers
	}

	for i := len(container_rows) - 2; i >= 0; i-- {
		line := strings.Split(string(container_rows[i]), "")
		var containers []string
		for i := 1; i < len(line); i += 4 {
			containers = append(containers, string(line[i]))
		}
		for j := 0; j < len(containers); j++ {
			if containers[j] != " " {
				container_row := containers_map[j+1]
				container_row = append(container_row, containers[j])
				containers_map[j+1] = container_row
			}
		}
	}

	instructions := strings.Split(string(instruction_set), "\n")

	// ------------------------------------------------------------------ //
	// Part 1
	// ------------------------------------------------------------------ //
	containers_map = cratemover_9000(containers_map, instructions)

	keys := make([]int, 0, len(containers_map))

	for k := range containers_map {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	fmt.Println("---Part 1---")
	for _, stack := range keys {
		if len(containers_map[stack]) > 0 {
			containers := containers_map[stack]
			top_container := containers[len(containers)-1]
			fmt.Printf("Stack %d: %s\n", stack, top_container)
		}
	}

	// ------------------------------------------------------------------ //
	// Part 2
	// ------------------------------------------------------------------ //

	// Rebuild Container Map
	for i := 1; i <= stack_count; i++ {
		var containers []string
		containers_map[i] = containers
	}

	for i := len(container_rows) - 2; i >= 0; i-- {
		line := strings.Split(string(container_rows[i]), "")
		var containers []string
		for i := 1; i < len(line); i += 4 {
			containers = append(containers, string(line[i]))
		}
		for j := 0; j < len(containers); j++ {
			if containers[j] != " " {
				container_row := containers_map[j+1]
				container_row = append(container_row, containers[j])
				containers_map[j+1] = container_row
			}
		}
	}

	containers_map = cratemover_9001(containers_map, instructions)

	fmt.Printf("\n---Part 2---\n")
	for _, stack := range keys {
		if len(containers_map[stack]) > 0 {
			containers := containers_map[stack]
			top_container := containers[len(containers)-1]
			fmt.Printf("Stack %d: %s\n", stack, top_container)
		}
	}
}
