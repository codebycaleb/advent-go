package solutions_test

import (
	"testing"

	"github.com/codebycaleb/advent-go/y2020/solutions"
)

const input1 = `1721
979
366
299
675
1456`

func TestDay1Part1(t *testing.T) {
	const want = "514579"
	result := solutions.Day1().Part1(input1)
	if result != want {
		t.Fatalf("Part 1 did not execute successfully (expected %v but got %v).", want, result)
	}
}

func TestDay1Part2(t *testing.T) {
	const want = "241861950"
	result := solutions.Day1().Part2(input1)
	if result != want {
		t.Fatalf("Part 2 did not execute successfully (expected %v but got %v).", want, result)
	}
}
