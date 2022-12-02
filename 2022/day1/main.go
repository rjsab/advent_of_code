package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var inventories = make(map[int]int)

	file, err := os.Open("../inputs/day1.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	elf_count := 1
	inventories[elf_count] = 0
	for scanner.Scan() {
		if scanner.Text() != "" {
			calories, _ := strconv.Atoi(scanner.Text())
			new_cal_count := inventories[elf_count] + calories

			inventories[elf_count] = new_cal_count
		} else if scanner.Text() == "" {
			elf_count++
			inventories[elf_count] = 0
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	keys := make([]int, 0, len(inventories))

	for key := range inventories {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return inventories[keys[i]] > inventories[keys[j]]
	})

	top_3_cals := 0
	for i := 0; i < 3; i++ {
		top_3_cals += inventories[keys[i]]
	}

	fmt.Printf("Max Calorie Elf: %d\n", inventories[keys[0]])
	fmt.Printf("Top 3 Calorie Total: %d\n", top_3_cals)
}
