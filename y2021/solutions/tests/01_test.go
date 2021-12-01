package solutions_test

import (
	"testing"

	"github.com/codebycaleb/advent-go/y2021/solutions"
)

const input01 = `199
200
208
210
200
207
240
269
260
263`

func TestDay01Part1(t *testing.T) {
	const want = "7"
	result := solutions.Day01().Part1(input01)
	if result != want {
		t.Fatalf("Part 1 did not execute successfully (expected %v but got %v).", want, result)
	}
}

func TestDay01Part2(t *testing.T) {
	const want = "5"
	result := solutions.Day01().Part2(input01)
	if result != want {
		t.Fatalf("Part 2 did not execute successfully (expected %v but got %v).", want, result)
	}
}
