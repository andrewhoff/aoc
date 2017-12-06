package part1

// AOC day 2, part 1

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
		min := getMin(rowCols)
		max := getMax(rowCols)
		diff := max - min
		checksum += diff
	}

	return checksum
}

func getMin(rowCol []string) int {
	val, _ := strconv.Atoi(rowCol[0])
	min := val

	for _, colVal := range rowCol {
		if v, err := strconv.Atoi(colVal); err == nil {
			if v < min {
				min = v
			}
		}
	}

	return min
}

func getMax(rowCol []string) int {
	val, _ := strconv.Atoi(rowCol[0])
	max := val

	for _, colVal := range rowCol {
		if v, err := strconv.Atoi(colVal); err == nil {
			if v > max {
				max = v
			}
		}
	}

	return max
}
