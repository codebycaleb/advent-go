package solutions_test

import (
	"testing"

	"github.com/codebycaleb/advent-go/y2020/solutions"
)

const input06 = `abc

a
b
c

ab
ac

a
a
a
a

b`

func TestDay06Part1(t *testing.T) {
	const want = "11"
	result := solutions.Day06().Part1(input06)
	if result != want {
		t.Fatalf("Part 1 did not execute successfully (expected %v but got %v).", want, result)
	}
}

func TestDay06Part2(t *testing.T) {
	const want = "6"
	result := solutions.Day06().Part2(input06)
	if result != want {
		t.Fatalf("Part 2 did not execute successfully (expected %v but got %v).", want, result)
	}
}
