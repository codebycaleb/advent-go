package solutions

import (
	"fmt"

	"github.com/codebycaleb/advent-go/y2020/utils"
)

func part1(entries []int) (int, int) {
	seen := map[int]struct{}{}
	for _, x := range entries {
		inverse_x := 2020 - x
		_, ok := seen[inverse_x]
		if ok {
			return inverse_x, x
		}
		seen[x] = struct{}{}
	}
	return -1, -1
}

type Sum struct {
	X int
	Y int
}

func part2(entries []int) (int, int, int) {
	sums := map[int]Sum{}
	for i, x := range entries {
		sum, ok := sums[2020-x]
		if ok {
			return x, sum.X, sum.Y
		}
		for _, y := range entries[i+1:] {
			if x+y < 2020 {
				sums[x+y] = Sum{x, y}
			}
		}
	}
	return -1, -1, -1
}

func Day1Part1(input string) string {
	entries := utils.Integers(input)
	x, y := part1(entries)
	return fmt.Sprint(x * y)
}

func Day1Part2(input string) string {
	entries := utils.Integers(input)
	x, y, z := part2(entries)
	return fmt.Sprint(x * y * z)
}

func Day1() utils.Day {
	return utils.Day{ID: "01", Part1: Day1Part1, Part2: Day1Part2}
}
