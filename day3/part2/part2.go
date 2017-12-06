package part2

// AOC day 3, part 2

import (
	"fmt"
	"strconv"
	"strings"
)

// Solve ...
func Solve(puzzInput string) int {
	fmt.Println("Solving checksum....")
	checksum := 0

	rows := strings.Split(puzzInput, "\n")

	for _, row := range rows {
		rowCols := strings.Split(row, "\t")
		result := getEquallyDivisibleResult(rowCols)
		checksum += result
	}

	return checksum
}

func getEquallyDivisibleResult(rowCol []string) int {
	rowColInts := make([]int, 0)
	result := 0

	for _, v := range rowCol {
		val, _ := strconv.Atoi(v)
		rowColInts = append(rowColInts, val)
	}

	for i := 0; i < len(rowColInts); i++ {
		if result = findDivisible(rowColInts[i], rowColInts); result > 0 {
			return result
		}
	}

	return result
}

func findDivisible(n int, arr []int) int {
	for _, v := range arr {
		if v == n {
			continue
		}
		if v%n == 0 {
			return v / n
		}
	}

	return 0
}
