package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {

	file, err := os.ReadFile("../test_inputs/day5.txt")
	expectedOutput := "CMZ"

	if err != nil {
		fmt.Println(err)
	}

	var containers_map = make(map[int][]string)
	containers_set := strings.Split(string(file), "\n\n")[0]
	instruction_set := strings.Split(string(file), "\n\n")[1]
	instructions := strings.Split(string(instruction_set), "\n")

	containers_map = generate_containers_map(containers_set, containers_map)
	containers_map = cratemover_9000(containers_map, instructions)
	output := generate_output(containers_map)

	if output != expectedOutput {
		t.Errorf("Expected: %s, Actual: %s", expectedOutput, output)
	}

}

func TestPart2(t *testing.T) {

	file, err := os.ReadFile("../test_inputs/day5.txt")
	expectedOutput := "MCD"

	if err != nil {
		fmt.Println(err)
	}

	var containers_map = make(map[int][]string)
	containers_set := strings.Split(string(file), "\n\n")[0]
	instruction_set := strings.Split(string(file), "\n\n")[1]
	instructions := strings.Split(string(instruction_set), "\n")

	containers_map = generate_containers_map(containers_set, containers_map)
	containers_map = cratemover_9001(containers_map, instructions)
	output := generate_output(containers_map)

	if output != expectedOutput {
		t.Errorf("Expected: %s, Actual: %s", expectedOutput, output)
	}

}
