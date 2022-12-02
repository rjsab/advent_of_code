package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	/* Game Rules:
			Rock beats Scissor
			Scissor beats Paper
		    Paper beats Rock

		    Rock =  A, X  worth 1 point
		    Paper =  B, Y worth 2 points
		    Scissor = C, Z worth 3 points

			Scoring:
		    Win = 6
		    Draw = 3
		    Loss = 0

	   	    Round Score =  Guess_Value + Result_Value
	*/

	var game = map[string]int{
		"A":    1,
		"B":    2,
		"C":    3,
		"X":    1,
		"Y":    2,
		"Z":    3,
		"WIN":  6,
		"LOSS": 0,
		"DRAW": 3,
	}

	var pt1_tournament_total int = 0
	var pt2_tournament_total int = 0

	file, err := os.Open("../inputs/day2.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var round_total int
		guesses := strings.Split(scanner.Text(), " ")
		my_guess := guesses[1]
		opp_guess := guesses[0]

		switch my_guess {
		case "X":
			switch opp_guess {
			case "A":
				round_total = game[my_guess] + game["DRAW"]
			case "B":
				round_total = game[my_guess] + game["LOSS"]
			case "C":
				round_total = game[my_guess] + game["WIN"]
			}
		case "Y":
			switch opp_guess {
			case "A":
				round_total = game[my_guess] + game["WIN"]
			case "B":
				round_total = game[my_guess] + game["DRAW"]
			case "C":
				round_total = game[my_guess] + game["LOSS"]
			}
		case "Z":
			switch opp_guess {
			case "A":
				round_total = game[my_guess] + game["LOSS"]
			case "B":
				round_total = game[my_guess] + game["WIN"]
			case "C":
				round_total = game[my_guess] + game["DRAW"]
			}
		}

		pt1_tournament_total += round_total

		result := guesses[1]

		/*Part 2
		X = Loss
		Y = Draw
		Z = WIN
		*/
		switch result {
		case "X":
			switch opp_guess {
			case "A":
				round_total = game["C"] + game["LOSS"]
			case "B":
				round_total = game["A"] + game["LOSS"]
			case "C":
				round_total = game["B"] + game["LOSS"]
			}
		case "Y":
			switch opp_guess {
			case "A":
				round_total = game["A"] + game["DRAW"]
			case "B":
				round_total = game["B"] + game["DRAW"]
			case "C":
				round_total = game["C"] + game["DRAW"]
			}
		case "Z":
			switch opp_guess {
			case "A":
				round_total = game["B"] + game["WIN"]
			case "B":
				round_total = game["C"] + game["WIN"]
			case "C":
				round_total = game["A"] + game["WIN"]
			}
		}

		pt2_tournament_total += round_total
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Part 1 Tournament Total: %d\n", pt1_tournament_total)
	fmt.Printf("Part 2 Tournament Total: %d\n", pt2_tournament_total)
}
