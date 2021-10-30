package solutions_test

import (
	"testing"

	"github.com/codebycaleb/advent-go/y2020/solutions"
)

const input02 = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

func TestDay02Part1(t *testing.T) {
	const want = "2"
	result := solutions.Day02().Part1(input02)
	if result != want {
		t.Fatalf("Part 1 did not execute successfully (expected %v but got %v).", want, result)
	}
}

func TestDay02Part2(t *testing.T) {
	const want = "1"
	result := solutions.Day02().Part2(input02)
	if result != want {
		t.Fatalf("Part 2 did not execute successfully (expected %v but got %v).", want, result)
	}
}
