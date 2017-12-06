package part1_test

import (
	"testing"

	taxicab "github.com/andrewhoff/aoc/day3/part1"
)

func TestSample(t *testing.T) {
	var puzzInput = 1
	var expected = 0

	if ans := taxicab.Solve(puzzInput); ans != expected {
		t.Fatalf("solve for %d doesn't give expected answer of %d, instead got %d", puzzInput, expected, ans)
	}
}

func TestSample2(t *testing.T) {
	var puzzInput = 12
	var expected = 3

	if ans := taxicab.Solve(puzzInput); ans != expected {
		t.Fatalf("solve for %d doesn't give expected answer of %d, instead got %d", puzzInput, expected, ans)
	}
}

func TestSample3(t *testing.T) {
	var puzzInput = 23
	var expected = 2

	if ans := taxicab.Solve(puzzInput); ans != expected {
		t.Fatalf("solve for %d doesn't give expected answer of %d, instead got %d", puzzInput, expected, ans)
	}
}

func TestSample4(t *testing.T) {
	var puzzInput = 1024
	var expected = 31

	if ans := taxicab.Solve(puzzInput); ans != expected {
		t.Fatalf("solve for %d doesn't give expected answer of %d, instead got %d", puzzInput, expected, ans)
	}
}

func TestMyInput(t *testing.T) {
	var puzzInput = 347991
	var expected = 1

	if ans := taxicab.Solve(puzzInput); ans != expected {
		t.Fatalf("solve for %d doesn't give expected answer of %d, instead got %d", puzzInput, expected, ans)
	}
}
