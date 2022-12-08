package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check_row_visibility(forest [][]int, row int, col int) bool {

	// Check Row visibility
	tree_row := forest[row]
	current_tree := forest[row][col]

	// Left Side
	var left_vis = []int{}
	for _, tree := range tree_row[:col] {
		if tree < current_tree {
			left_vis = append(left_vis, tree_row[col])
		}
	}
	// Right Side
	var right_vis = []int{}
	for _, tree := range tree_row[col+1:] {
		if tree < current_tree {
			right_vis = append(right_vis, tree_row[col])
		}
	}

	if len(left_vis) == len(forest[:col]) || (len(right_vis) == len(forest[col+1:])) {
		return true
	}
	return false
}

func check_col_visibility(forest [][]int, row int, col int) bool {
	// Check Col visibility
	current_tree := forest[row][col]

	// Top Side
	var top_vis = []int{}
	for _, tree_row := range forest[:row] {
		if tree_row[col] < current_tree {
			top_vis = append(top_vis, tree_row[col])
		}
	}
	// Bottom Side
	var bot_vis = []int{}
	for _, tree_row := range forest[row+1:] {
		if tree_row[col] < current_tree {
			bot_vis = append(bot_vis, tree_row[col])
		}
	}

	if len(top_vis) == len(forest[:row]) || (len(bot_vis) == len(forest[row+1:])) {
		return true
	}

	return false
}

func check_view_score(forest [][]int) (scores []int) {
	/*
		Starting Tree since all perimeter trees
		will have a score of 0.
	*/
	row := 1
	col := 1

	for row < len(forest)-1 {
		for col < len(forest[row])-1 {
			// UP
			up_total := 0
			for i := row - 1; i >= 0; i-- {
				if forest[i][col] < forest[row][col] {
					up_total++
				} else if forest[i][col] >= forest[row][col] {
					up_total++
					break
				}
			}
			// Down
			down_total := 0
			for i := row + 1; i < len(forest); i++ {
				if forest[i][col] < forest[row][col] {
					down_total++
				} else if forest[i][col] >= forest[row][col] {
					down_total++
					break
				}
			}
			// Left
			left_total := 0
			for i := col - 1; i >= 0; i-- {
				if forest[row][i] < forest[row][col] {
					left_total++
				} else if forest[row][i] >= forest[row][col] {
					left_total++
					break
				}
			}
			// Right
			right_total := 0
			for i := col + 1; i < len(forest[row]); i++ {
				if forest[row][i] < forest[row][col] {
					right_total++
				} else if forest[row][i] >= forest[row][col] {
					right_total++
					break
				}
			}

			score := (up_total * down_total * left_total * right_total)
			scores = append(scores, score)
			col++
		}
		row++
		col = 1
	}

	return
}

func parse_input(file *os.File) (forest [][]int) {
	scanner := bufio.NewScanner(file)

	row := 0
	for scanner.Scan() {
		tree_row := strings.Split(scanner.Text(), "")

		var tree_column = []int{}

		for _, tree := range tree_row {
			tree_height, _ := strconv.Atoi(tree)
			tree_column = append(tree_column, tree_height)
		}

		forest = append(forest, tree_column)
		row++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return
}

func count_trees(forest [][]int) (count int) {

	count = (len(forest) * 2) + ((len(forest) - 2) * 2)

	for row := 1; row < len(forest)-1; row++ {
		for col := 1; col < (len(forest[row]) - 1); col++ {
			row_vis := check_row_visibility(forest, row, col)
			col_vis := check_col_visibility(forest, row, col)

			if row_vis || col_vis {
				count++
			}
		}
	}

	return
}

func find_max(scores []int) (max int) {
	max = scores[0]
	for _, score := range scores {
		if score > max {
			max = score
		}
	}
	return
}

func main() {
	file, err := os.Open("../inputs/day8.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	forest := parse_input(file)
	fmt.Println(count_trees(forest))
	scores := check_view_score(forest)
	fmt.Println(find_max(scores))
}
