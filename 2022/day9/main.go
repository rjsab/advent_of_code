package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Head struct {
	x int
	y int
}

type Tail struct {
	x int
	y int
}

func draw_board(head *Head, tail *Tail, gameboard *[][]string, direction string, spaces int) {
	switch direction {
	case "U":
		rows := head.y - spaces
		if rows < 0 {
			draw_cols := []string{}
			for x := 0; x < len((*gameboard)[head.x]); x++ {
				draw_cols = append(draw_cols, "+")
			}
			for y := 0; y < int(math.Abs(float64(rows))); y++ {
				*gameboard = append([][]string{draw_cols}, *gameboard...)
			}
			head.y = head.y + int(math.Abs(float64(rows)))
			tail.y = tail.y + int(math.Abs(float64(rows)))
		}
	case "D":
		rows := (head.y + spaces) - (len(*gameboard) - 1)
		if rows > 0 {
			draw_cols := []string{}
			for x := 0; x < len((*gameboard)[0]); x++ {
				draw_cols = append(draw_cols, "+")
			}
			for y := 0; y < rows; y++ {
				*gameboard = append(*gameboard, [][]string{draw_cols}...)
			}
		}
	case "L":
		cols := (head.x - spaces)
		if cols < 0 {
			draw_cols := []string{}
			for x := 0; x < int(math.Abs(float64(cols))); x++ {
				draw_cols = append(draw_cols, "+")
			}
			for y := 0; y < len(*(gameboard)); y++ {
				(*gameboard)[y] = append(draw_cols, (*gameboard)[y]...)
			}
			head.x = head.x + int(math.Abs(float64(cols)))
			tail.x = tail.x + int(math.Abs(float64(cols)))
		}
	case "R":
		cols := (head.x + spaces) - (len((*gameboard)[head.x]) - 1)
		if cols > 0 {
			draw_cols := []string{}
			for x := 0; x < cols; x++ {
				draw_cols = append(draw_cols, "+")
			}
			for y := 0; y < len(*(gameboard)); y++ {
				(*gameboard)[y] = append((*gameboard)[y], draw_cols...)
			}
		}
	}
}

func move_head(direction string, head *Head, gameboard *[][]string) {

	switch direction {
	case "U":
	case "D":
	case "L":
	case "R":
	}
}

func main() {
	file, err := os.Open("../test_inputs/day9.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	head := Head{
		x: 0,
		y: 0,
	}

	tail := Tail{
		x: 0,
		y: 0,
	}

	//gameboard[y][x]
	gameboard := [][]string{{"#"}}

	for scanner.Scan() {
		action := strings.Split(scanner.Text(), " ")
		direction := strings.ToUpper(action[0])
		spaces, _ := strconv.Atoi(action[1])

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}

		fmt.Println(direction, spaces)
		draw_board(&head, &tail, &gameboard, direction, spaces)

		for move := 0; move < spaces; move++ {
			/*
				TODO:
				Build Movement
			*/
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	for _, y := range gameboard {
		for _, x := range y {
			fmt.Printf("%s", x)
		}
		fmt.Println()
	}
}
