package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	name    int
	items   []int
	sign    string
	amount  string
	test    int
	iftrue  int
	iffalse int
	inspect int
}

func main() {
	//file, err := os.ReadFile("../test_inputs/day11.txt")
	file, err := os.ReadFile("../inputs/day11.txt")

	if err != nil {
		fmt.Println(err)
	}

	monkey_file := strings.Split(string(file), "\n\n")

	var monkeys = []Monkey{}

	for _, monkey := range monkey_file {

		instructions := strings.Split(monkey, "\n")

		// Parse Monkey ID
		monkey := strings.Split(instructions[0], " ")[1]
		monkey = strings.TrimSuffix(monkey, ":")
		monkey_name, _ := strconv.Atoi(monkey)

		// Parse Item List
		var item_list = []int{}
		items := strings.SplitAfter(instructions[1], ":")[1]
		items_list := strings.Split(strings.TrimSpace(items), ",")
		for _, item := range items_list {
			item_i, _ := strconv.Atoi(strings.TrimSpace(item))
			item_list = append(item_list, item_i)
		}

		// Parse Operation and Amount
		operation := strings.SplitAfter(instructions[2], ":")[1]
		sign := strings.Split(operation, " ")[4]
		amount := strings.Split(operation, " ")[5]

		// Parse Test
		test := strings.SplitAfter(instructions[3], ":")[1]
		test_amount, _ := strconv.Atoi(strings.Split(test, " ")[3])

		// Parse True
		iftrue := strings.SplitAfter(instructions[4], ":")[1]
		true_monkey, _ := strconv.Atoi(strings.Split(iftrue, " ")[4])

		// Parse False
		iffalse := strings.SplitAfter(instructions[5], ":")[1]
		false_monkey, _ := strconv.Atoi(strings.Split(iffalse, " ")[4])

		monkeys = append(monkeys, Monkey{
			name:    monkey_name,
			items:   item_list,
			sign:    sign,
			amount:  amount,
			test:    test_amount,
			iftrue:  true_monkey,
			iffalse: false_monkey,
			inspect: 0,
		})
	}

	// Part 1
	//var rounds int = 20

	//Part 2
	var rounds int = 10000

	common_mult := 1
	for _, m := range monkeys {
		common_mult *= m.test
	}

	for i := 1; i <= rounds; i++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				var amount int
				if monkey.amount == "old" {
					amount = item
				} else {
					amount, _ = strconv.Atoi(monkey.amount)
				}
				switch monkey.sign {
				case "*":
					item = item * amount
				case "+":
					item = item + amount
				}

				// Part 1
				//item = item / 3
				// Part 2
				item = item % common_mult

				if item%monkey.test == 0 {
					//fmt.Printf("Item with worry %d Pass to monkey %d\n", item, monkey.iftrue)
					monkeys[monkey.iftrue].items = append(monkeys[monkey.iftrue].items, item)
				} else {
					//fmt.Printf("Item with worry %d Pass to monkey %d\n", item, monkey.iffalse)
					monkeys[monkey.iffalse].items = append(monkeys[monkey.iffalse].items, item)
				}
				monkeys[i].items = monkeys[i].items[1:]
				monkeys[i].inspect = monkeys[i].inspect + 1

			}
		}

	}
	top_1 := 0
	top_2 := 0
	for _, monkey := range monkeys {
		if monkey.inspect > top_1 {
			top_2 = top_1
			top_1 = monkey.inspect
		} else if monkey.inspect > top_2 {
			top_2 = monkey.inspect
		}
		fmt.Printf("Monkey %d inspected items %d times.\n", monkey.name, monkey.inspect)
	}
	fmt.Println(top_1 * top_2)
}
