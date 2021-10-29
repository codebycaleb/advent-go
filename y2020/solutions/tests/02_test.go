package solutions_test

import (
	"testing"

	"github.com/codebycaleb/advent-go/y2020/solutions"
)

const input2 = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

func TestDay2Part1(t *testing.T) {
	const want = "2"
	result := solutions.Day2().Part1(input2)
	if result != want {
		t.Fatalf("Part 1 did not execute successfully (expected %v but got %v).", want, result)
	}
}

func TestDay2Part2(t *testing.T) {
	const want = "1"
	result := solutions.Day2().Part2(input2)
	if result != want {
		t.Fatalf("Part 2 did not execute successfully (expected %v but got %v).", want, result)
	}
}
