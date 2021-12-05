package solutions_test

import (
	"testing"

	"github.com/codebycaleb/advent-go/y2021/solutions"
)

const input03 = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

func TestDay03Part1(t *testing.T) {
	const want = "198"
	result := solutions.Day03().Part1(input03)
	if result != want {
		t.Fatalf("Part 1 did not execute successfully (expected %v but got %v).", want, result)
	}
}

func TestDay03Part2(t *testing.T) {
	const want = "230"
	result := solutions.Day03().Part2(input03)
	if result != want {
		t.Fatalf("Part 2 did not execute successfully (expected %v but got %v).", want, result)
	}
}
