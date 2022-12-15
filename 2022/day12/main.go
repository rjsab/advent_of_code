package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var alphabet = make(map[string]int)

	index := 1
	for character := 'a'; character <= 'z'; character++ {
		alphabet[string(character)] = index
		index++
	}

	file, err := os.ReadFile("../test_inputs/day12.txt")

	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(file), "\n")

	var path = []string{}
	start_index := 0
	end_index := 0
	vindex := len(lines[0])

	for _, line := range lines {
		path = append(path, strings.Split(line, "")...)

	}

	for i, char := range path {
		if char == "S" {
			start_index = i
		} else if char == "E" {
			end_index = i
		}
	}

	for i := 0; i < len(path); i += vindex {
		fmt.Println(path[i : i+vindex])
	}

}
