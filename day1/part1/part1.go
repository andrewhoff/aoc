package part1

// AOC day 1, part 1

import (
	"fmt"
	"strconv"
)

// Solve ...
func Solve(ints []int) int {
	fmt.Println("Solving CAPTCHA....")
	sum := 0
	temp := sum

	for i := 0; i < len(ints); i++ {
		if ints[i] == temp {
			sum += ints[i]
		}

		temp = ints[i]
	}

	if ints[0] == ints[len(ints)-1] {
		sum += ints[0]
	}

	return sum
}

// IntsListFromInput ...
func IntsListFromInput(puzzInput string) []int {
	list := make([]int, 0)
	for i := 0; i < len(puzzInput); i++ {
		if val, err := strconv.Atoi(string(puzzInput[i])); err == nil {
			list = append(list, val)
		}
	}

	return list
}
