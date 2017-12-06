package part1

// AOC day 3, part 1

import (
	"fmt"
	"log"
	"math"
)

// Solve ...
func Solve(puzzInput int) int {
	fmt.Println("Solving taxicab distance....")

	var currBottomRightCorner float64
	currBottomRightCorner = 1
	var sideLength float64
	sideLength = 1
	var steps float64

	for {
		if float64(puzzInput) > currBottomRightCorner {
			steps = (math.Sqrt(currBottomRightCorner) + 1)
			sideLength = (math.Sqrt(currBottomRightCorner) + 2)
			fmt.Println(steps)
			currBottomRightCorner = math.Pow((math.Sqrt(currBottomRightCorner) + 2), 2)
			// fmt.Printf("New current bottom right corner: %d\n", int(currBottomRightCorner))
		} else {
			log.Println("Breaking....")
			log.Printf("steps: %d\n", int(steps))
			log.Printf("length of sides: %d\n", int(sideLength))
			log.Printf("Diff: %d", int(currBottomRightCorner)-puzzInput)
			diff := int(currBottomRightCorner) - puzzInput
			steps -= (sideLength - float64(diff%int(sideLength)))
			break
		}
	}

	return int(steps)
}
