package solutions_test

import (
	"testing"

	"github.com/codebycaleb/advent-go/y2021/solutions"
)

const input02 = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func TestDay02Part1(t *testing.T) {
	const want = "150"
	result := solutions.Day02().Part1(input02)
	if result != want {
		t.Fatalf("Part 1 did not execute successfully (expected %v but got %v).", want, result)
	}
}

func TestDay02Part2(t *testing.T) {
	const want = "900"
	result := solutions.Day02().Part2(input02)
	if result != want {
		t.Fatalf("Part 2 did not execute successfully (expected %v but got %v).", want, result)
	}
}
