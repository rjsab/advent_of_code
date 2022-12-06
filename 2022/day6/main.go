package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func stream_scanner(line string, delimeter int) int {

	message_length := delimeter
	line_split := strings.Split(line, "")

	for delimeter <= len(line_split) {

		var stream = make(map[string]bool)

		for _, v := range line_split[delimeter-message_length : delimeter] {
			stream[v] = true
		}

		if len(stream) == message_length {
			break
		}
		delimeter++
	}

	return delimeter
}

func main() {
	file, err := os.Open("../inputs/day6.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var count int

	for scanner.Scan() {
		// Part 1
		count = stream_scanner(scanner.Text(), 4)

		fmt.Printf("Stream Count: %d\n", count)

		// Part 2
		count = stream_scanner(scanner.Text(), 14)

		fmt.Printf("Stream Count: %d\n", count)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
