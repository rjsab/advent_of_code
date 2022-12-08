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

	var pdirs = []string{}
	dir_sizes = dir_traversal(&root, pdirs, 0, dir_sizes)

	actual := calculate_total_size(dir_sizes)

	if actual != expectedOutput {
		t.Errorf("Expected: %d, Actual: %d", expectedOutput, actual)
	}
}

func TestPart2(t *testing.T) {

	file, err := os.ReadFile("../test_inputs/day7.txt")
	if err != nil {
		fmt.Println(err)
	}

	var expectedOutput int = 24933642
	var dir_sizes = make(map[string]int)
	lines := strings.Split(string(file), "\n")
	root := command_parse(lines)

	var pdirs = []string{}
	dir_sizes = dir_traversal(&root, pdirs, 0, dir_sizes)

	calculate_total_size(dir_sizes)
	actual := (download_filespace(70000000, 30000000, dir_sizes))

	if actual != expectedOutput {
		t.Errorf("Expected: %d, Actual: %d", expectedOutput, actual)
	}
}
