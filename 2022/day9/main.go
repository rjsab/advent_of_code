package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Piece struct {
	x int
	y int
}

func calculate_tail_pos(gameboard *[][]string) (count int) {
	for _, y := range *gameboard {
		for _, x := range y {
			if x == "*" {
				count++
			}
		}
	}
	return
}

func draw_board(pieces *[]Piece, gameboard *[][]string, direction string, spaces int) {
	head := (*pieces)[0]
	switch direction {
	case "U":
		rows := head.y - spaces
		if rows < 0 {
			new_rows := [][]string{}
			for y := 0; y < int(math.Abs(float64(rows))); y++ {
				new_cols := []string{}
				for x := 0; x < len((*gameboard)[head.y]); x++ {
					new_cols = append(new_cols, ".")
				}
				new_rows = append(new_rows, new_cols)
			}

			*gameboard = append(new_rows, (*gameboard)...)

			for i := range *pieces {
				(*pieces)[i].y = (*pieces)[i].y + int(math.Abs(float64(rows)))
			}
		}
	case "D":
		rows := (head.y + spaces) - (len(*gameboard) - 1)
		if rows > 0 {
			new_rows := [][]string{}
			for y := 0; y < rows; y++ {
				new_cols := []string{}
				for x := 0; x < len((*gameboard)[head.y]); x++ {
					new_cols = append(new_cols, ".")
				}
				new_rows = append(new_rows, new_cols)
			}
			*gameboard = append((*gameboard), new_rows...)
		}
	case "L":
		cols := (head.x - spaces)
		if cols < 0 {
			for i := range *gameboard {
				new_cols := []string{}
				for x := 0; x < int(math.Abs(float64(cols))); x++ {
					new_cols = append(new_cols, ".")
				}
				(*gameboard)[i] = append(new_cols, (*gameboard)[i]...)
			}
			for i := range *pieces {
				(*pieces)[i].x = (*pieces)[i].x + int(math.Abs(float64(cols)))
			}
		}
	case "R":
		cols := (head.x + spaces) - (len((*gameboard)[head.y]) - 1)
		if cols > 0 {
			for i := range *gameboard {
				new_cols := []string{}
				for x := 0; x < cols; x++ {
					new_cols = append(new_cols, ".")
				}
				(*gameboard)[i] = append((*gameboard)[i], new_cols...)
			}
		}
	}
	return
}

func play_game(file *os.File, pieces int) (gameboard [][]string) {

	scanner := bufio.NewScanner(file)

	var game_pieces = []Piece{}

	for piece := 0; piece < pieces; piece++ {
		game_pieces = append(game_pieces, Piece{x: 0, y: 0})
	}

	//gameboard[y][x]
	gameboard = [][]string{{"s"}}

	for scanner.Scan() {
		action := strings.Split(scanner.Text(), " ")
		direction := strings.ToUpper(action[0])
		spaces, _ := strconv.Atoi(action[1])

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}

		draw_board(&game_pieces, &gameboard, direction, spaces)

		for move := 0; move < spaces; move++ {
			tail := &game_pieces[len(game_pieces)-1]
			switch direction {
			case "U":
				game_pieces[0].y--
			case "D":
				game_pieces[0].y++
			case "L":
				game_pieces[0].x--
			case "R":
				game_pieces[0].x++
			}

			for i := 1; i < len(game_pieces); i++ {

				x_delta := game_pieces[i-1].x - game_pieces[i].x
				y_delta := game_pieces[i-1].y - game_pieces[i].y

				if math.Abs(float64(x_delta)) <= 1 && math.Abs(float64(y_delta)) <= 1 {
					break
				}
				if x_delta > 0 {
					game_pieces[i].x++
				} else if x_delta < 0 {
					game_pieces[i].x--
				}
				if y_delta > 0 {
					game_pieces[i].y++
				} else if y_delta < 0 {
					game_pieces[i].y--
				}
			}
			gameboard[tail.y][tail.x] = "*"
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return
}

func main() {
	file, err := os.Open("../inputs/day9.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	gameboard := play_game(file, 10)

	fmt.Println()
	for _, y := range gameboard {
		for _, x := range y {
			fmt.Printf("%s", x)
		}
		fmt.Println()
	}

	fmt.Printf("Unique Tail Pos: %d\n", calculate_tail_pos(&gameboard))
}
