package solutions

import (
	"sort"
	"strconv"

	"github.com/codebycaleb/advent-go/common"
	"github.com/codebycaleb/advent-go/common/utils"
)

func Day05() common.Day {

	type Pass struct {
		row, col, id int
	}

	// A seat might be specified like FBFBBFFRLR, where F means "front", B means "back", L means "left", and R means "right".
	// The first 7 characters represent the row in binary, where F = 0 and B = 1.
	// The last 3 characters represent the column in binary, where L = 0 and R = 1.
	// Finally, the ID is calculated by multiplying the row by 8, then adding the column.
	parsePass := func(pass string) Pass {
		row := 0
		col := 0
		for i := 0; i < 7; i++ {
			if pass[i] == 'B' {
				row |= 1 << (6 - i)
			}
		}
		for i := 0; i < 3; i++ {
			if pass[7+i] == 'R' {
				col |= 1 << (2 - i)
			}
		}
		return Pass{row, col, row*8 + col}
	}

	part1 := func(input string) string {
		lines := utils.Lines(input)
		// parse each line
		passes := make([]Pass, len(lines))
		for i, line := range lines {
			passes[i] = parsePass(line)
		}
		// find max
		max := 0
		for _, pass := range passes {
			if pass.id > max {
				max = pass.id
			}
		}
		return strconv.Itoa(max)
	}

	part2 := func(input string) string {
		lines := utils.Lines(input)
		// parse each line
		passes := make([]Pass, len(lines))
		for i, line := range lines {
			passes[i] = parsePass(line)
		}
		// sort by .id
		sort.Slice(passes, func(i, j int) bool {
			return passes[i].id < passes[j].id
		})
		// find first skipped .id value
		for i := 1; i < len(passes); i++ {
			if passes[i].id != passes[i-1].id+1 {
				return strconv.Itoa(passes[i-1].id + 1)
			}
		}
		return "-1"
	}

	return common.Day{ID: "05", Part1: part1, Part2: part2}
}
