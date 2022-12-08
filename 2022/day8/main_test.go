package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {

	file, err := os.Open("../test_inputs/day8.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	forest := parse_input(file)
	actual := count_trees(forest)
	expectedOutput := 21

	if actual != expectedOutput {
		t.Errorf("Expected: %d, Actual: %d", expectedOutput, actual)
	}
}

func TestPart2(t *testing.T) {

	file, err := os.Open("../test_inputs/day8.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	forest := parse_input(file)
	scores := check_view_score(forest)
	actual := find_max(scores)

	expectedOutput := 8

	if actual != expectedOutput {
		t.Errorf("Expected: %d, Actual: %d", expectedOutput, actual)
	}
}
