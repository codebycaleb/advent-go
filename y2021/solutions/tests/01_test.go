package solutions_test

import (
	"testing"

	"github.com/codebycaleb/advent-go/y2020/solutions"
)

const input01 = `1721
979
366
299
675
1456`

func TestDay01Part1(t *testing.T) {
	const want = "514579"
	result := solutions.Day01().Part1(input01)
	if result != want {
		t.Fatalf("Part 1 did not execute successfully (expected %v but got %v).", want, result)
	}
}

func TestDay01Part2(t *testing.T) {
	const want = "241861950"
	result := solutions.Day01().Part2(input01)
	if result != want {
		t.Fatalf("Part 2 did not execute successfully (expected %v but got %v).", want, result)
	}
}
