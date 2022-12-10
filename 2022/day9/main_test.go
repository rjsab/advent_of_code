package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {

	file, err := os.Open("../test_inputs/day9.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		actual := follow_tail(scanner.Text(), game_board)
	}
	expectedOutput := 13

	if actual != expectedOutput {
		t.Errorf("Expected: %d, Actual: %d", expectedOutput, actual)
	}
}
