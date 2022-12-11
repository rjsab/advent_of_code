package main

import (
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

	gameboard := play_game(file, 10)

	fmt.Println()
	for _, y := range gameboard {
		for _, x := range y {
			fmt.Printf("%s  ", x)
		}
		fmt.Println()
	}

	actual := calculate_tail_pos(&gameboard)
	expectedOutput := 36

	if actual != expectedOutput {
		t.Errorf("Expected: %d, Actual: %d", expectedOutput, actual)
	}
}
