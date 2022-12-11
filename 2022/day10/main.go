package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func draw_pixel(sprite_index int, cycle int) string {
	if sprite_index == cycle-1 || sprite_index-1 == cycle-1 || sprite_index+1 == cycle-1 {
		return "#"
	} else {
		return "."
	}
}

func run_clock(file *os.File) {

	scanner := bufio.NewScanner(file)

	var screen = [6][]string{{}}
	var screen_row = 0
	var row = []string{}
	sprite_index := 1
	interval := 40
	cycle := 1

	for scanner.Scan() {
		action := strings.Split(scanner.Text(), " ")

		if action[0] == "addx" {
			count, _ := strconv.Atoi(action[1])
			row = append(row, draw_pixel(sprite_index, cycle))

			if cycle == interval {
				cycle = 0
				screen[screen_row] = row
				screen_row++
				row = []string{}
			}
			cycle++

			row = append(row, draw_pixel(sprite_index, cycle))
			if cycle == interval {
				cycle = 0
				screen[screen_row] = row
				screen_row++
				row = []string{}
			}
			sprite_index += count
			cycle++

		} else {
			row = append(row, draw_pixel(sprite_index, cycle))
			if cycle == interval {
				cycle = 0
				screen[screen_row] = row
				screen_row++
				row = []string{}
			}
			cycle++
			continue
		}
	}

	for _, row := range screen {
		fmt.Println(row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}

func main() {
	file, err := os.Open("../inputs/day10.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	run_clock(file)
}
