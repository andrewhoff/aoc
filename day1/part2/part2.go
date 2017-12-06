package part2

// AOC day 1, part 2

import (
	"fmt"
	"strconv"
)

// Solve ...
func Solve(ints []int) int {
	fmt.Println("Solving CAPTCHA....")
	sum := 0
	steps := len(ints) / 2

	for i := 0; i < len(ints); i++ {
		if ints[i] == ints[(i+steps)%len(ints)] {
			sum += ints[i]
		}
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
