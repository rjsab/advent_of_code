package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func intersection(arr1, arr2 []string) (intersect []string) {
	m := make(map[string]bool)

	for _, item := range arr1 {
		m[item] = true
	}

	for _, item := range arr2 {
		if _, ok := m[item]; ok {
			intersect = append(intersect, item)
			break
		}
	}
	return
}

func triple_intersect(string1, string2, string3 string) (intersect []string) {
	m1 := make(map[string]bool)
	m2 := make(map[string]bool)

	for _, item := range string1 {
		m1[string(item)] = true
	}

	for _, item := range string2 {
		if _, ok := m1[string(item)]; ok {
			m2[string(item)] = true
		}
	}

	for _, item := range string3 {
		if _, ok := m2[string(item)]; ok {
			intersect = append(intersect, string(item))
			break
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
		letter := intersection(pocket1, pocket2)
		priority_total += priority[letter[0]]

	}

	fmt.Printf("Priority Total: %d\n", priority_total)

	var elf_groups [][]string
	for i := 0; i < len(rucksacks); i += 3 {
		elf_groups = append(elf_groups, rucksacks[i:i+3])
	}

	group_priority_total := 0
	for _, group := range elf_groups {
		letter := triple_intersect(group[0], group[1], group[2])
		group_priority_total += priority[letter[0]]
	}

	fmt.Printf("Group Badge Priority Total: %d\n", group_priority_total)

}
