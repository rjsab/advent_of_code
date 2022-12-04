package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check_containment(a, b []int) (contains bool) {

	if (a[0] >= b[0]) && (a[1] <= b[1]) {
		contains = true
	} else if (a[0] <= b[0]) && (a[1] >= b[1]) {
		contains = true
	}

	return
}

func check_overlap(a, b []int) (overlap bool) {

	if (a[0] >= b[0]) && (a[0] <= b[1]) {
		overlap = true
	} else if (b[0] >= a[0]) && (b[0] <= a[1]) {
		overlap = true
	}

	return
}

func sanitize_section(section_range string) (section_list []int) {

	sections := strings.Split(section_range, "-")
	section_start, _ := strconv.Atoi(sections[0])
	section_end, _ := strconv.Atoi(sections[1])

	section_list = append(section_list, section_start)
	section_list = append(section_list, section_end)

	return
}

func main() {
	file, err := os.Open("../inputs/day4.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	full_contain_count := 0
	overlap_count := 0
	for scanner.Scan() {
		group_assignments := strings.Split(scanner.Text(), ",")

		var sanitized_section_range [][]int
		for _, section_range := range group_assignments {
			sanitized_section_range = append(sanitized_section_range, sanitize_section(section_range))
		}

		if check_containment(sanitized_section_range[0], sanitized_section_range[1]) {
			full_contain_count++
		}
		if check_overlap(sanitized_section_range[0], sanitized_section_range[1]) {
			overlap_count++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Full Contain Count: %d \n", full_contain_count)
	fmt.Printf("Overlap Count: %d \n", overlap_count)
}
