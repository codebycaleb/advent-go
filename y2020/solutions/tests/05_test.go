package solutions_test

import (
	"testing"

	"github.com/codebycaleb/advent-go/y2020/solutions"
)

const input05 = `FFFFFFFLLL
FFFFFFFLRL`

func TestDay05Part1(t *testing.T) {
	const want = "2"
	result := solutions.Day05().Part1(input05)
	if result != want {
		t.Fatalf("Part 1 did not execute successfully (expected %v but got %v).", want, result)
	}
}

func TestDay05Part2(t *testing.T) {
	const want = "1"
	result := solutions.Day05().Part2(input05)
	if result != want {
		t.Fatalf("Part 2 did not execute successfully (expected %v but got %v).", want, result)
	}
}
