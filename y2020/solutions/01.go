package solutions

import (
	"strconv"

	"github.com/codebycaleb/advent-go/common"
	"github.com/codebycaleb/advent-go/common/utils"
)

func Day01() common.Day {

	part1 := func(input string) string {
		entries := utils.Integers(input)
		seen := map[int]struct{}{}
		for _, x := range entries {
			inverse_x := 2020 - x
			_, ok := seen[inverse_x]
			if ok {
				return strconv.Itoa(inverse_x * x)
			}
			seen[x] = struct{}{}
		}
		return "err"
	}

	type Sum struct {
		X int
		Y int
	}

	part2 := func(input string) string {
		entries := utils.Integers(input)
		sums := map[int]Sum{}
		for i, x := range entries {
			sum, ok := sums[2020-x]
			if ok {
				return strconv.Itoa(x * sum.X * sum.Y)
			}
			for _, y := range entries[i+1:] {
				if x+y < 2020 {
					sums[x+y] = Sum{x, y}
				}
			}
		}
		return "err"
	}

	return common.Day{ID: "01", Part1: part1, Part2: part2}
}
