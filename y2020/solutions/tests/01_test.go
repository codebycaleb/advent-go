package solutions_test

import (
	"testing"

	"github.com/codebycaleb/advent-go/y2020/solutions"
)

const input = `1721
979
366
299
675
1456`

func TestPart1(t *testing.T) {
	const want = "514579"
	result := solutions.Day1Part1(input)
	if result != want {
		t.Fatalf("Part 1 did not execute successfully (expected %v but got %v).", want, result)
	}
}

func TestPart2(t *testing.T) {
	const want = "241861950"
	result := solutions.Day1Part2(input)
	if result != want {
		t.Fatalf("Part 2 did not execute successfully (expected %v but got %v).", want, result)
	}
}
