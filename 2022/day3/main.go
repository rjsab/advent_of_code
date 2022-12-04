package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func intersection(foo, bar []string) (intersect []string) {
	m := make(map[string]bool)

	for _, item := range foo {
		m[string(item)] = true
	}

	for _, item := range bar {
		if _, ok := m[string(item)]; ok {
			intersect = append(intersect, string(item))
		}
	}
	return
}

func main() {
	var priority = make(map[string]int)
	var rucksacks []string

	index := 1
	for character := 'a'; character <= 'z'; character++ {
		priority[string(character)] = index
		priority[string(unicode.ToUpper(character))] = index + 26
		index++
	}

	file, err := os.Open("../inputs/day3.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	priority_total := 0
	for scanner.Scan() {
		rucksack := scanner.Text()
		rucksacks = append(rucksacks, rucksack)
		sacks := strings.Split(rucksack, "")
		limit := len(rucksack) / 2

		pocket1 := sacks[0:limit]
		pocket2 := sacks[limit:len(rucksack)]
		letter := string(intersection(pocket1, pocket2)[0])
		priority_total += priority[letter]

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Priority Total: %d\n", priority_total)

	var elf_groups [][]string
	for i := 0; i < len(rucksacks); i += 3 {
		elf_groups = append(elf_groups, rucksacks[i:i+3])
	}

	group_priority_total := 0
	for _, rucksack := range elf_groups {
		rucksack1 := strings.Split(rucksack[0], "")
		rucksack2 := strings.Split(rucksack[1], "")
		rucksack3 := strings.Split(rucksack[2], "")
		letter := string(intersection(intersection(rucksack1, rucksack2), rucksack3)[0])
		group_priority_total += priority[letter]
	}

	fmt.Printf("Group Badge Priority Total: %d\n", group_priority_total)

}
