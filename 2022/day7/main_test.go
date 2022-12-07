package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {

	file, err := os.ReadFile("../test_inputs/day7.txt")
	if err != nil {
		fmt.Println(err)
	}

	var expectedOutput int = 95437
	var dir_sizes = make(map[string]int)
	lines := strings.Split(string(file), "\n")
	root := command_parse(lines)
	dir_sizes = dir_traversal(&root, 0, dir_sizes)

	actual := calculate_total_size(dir_sizes)

	if actual != expectedOutput {
		t.Errorf("Expected: %d, Actual: %d", expectedOutput, actual)
	}
}

// func TestPart2(t *testing.T) {}
/*
func TestPart2(t *testing.T) {

	file, err := os.Open("../test_inputs/day6.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var expectedOutput = []int{19, 23, 23, 29, 26}
	var delimeter int = 14
	var i = 0
	var count int
	for scanner.Scan() {
		count = stream_scanner(scanner.Text(), delimeter)

		if count != expectedOutput[i] {
			t.Errorf("Expected: %d, Actual: %d", expectedOutput[i], count)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
*/
