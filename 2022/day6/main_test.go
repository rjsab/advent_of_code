package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {

	file, err := os.Open("../test_inputs/day6.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var expectedOutput = []int{7, 5, 6, 10, 11}
	var delimeter int = 4
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

// func TestPart2(t *testing.T) {}
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
