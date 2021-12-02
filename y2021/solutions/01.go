package solutions

import (
	"strconv"

	"github.com/codebycaleb/advent-go/common"
	"github.com/codebycaleb/advent-go/common/utils"
)

func Day01() common.Day {

	part1 := func(input string) string {
		entries := utils.Integers(input)
		// count how many entries are larger than their previous entry
		var sum int
		for i := 0; i < len(entries)-1; i++ {
			if entries[i+1] > entries[i] {
				sum += 1
			}
		}
		return strconv.Itoa(sum)
	}

	part2 := func(input string) string {
		entries := utils.Integers(input)
		// map entry at index i to the sum of entries[i], entries[i+1], entries[i+2] as long as entries[i+2] exists
		sliding_windows := make([]int, len(entries)-2)
		for i := 0; i < len(entries)-2; i++ {
			sliding_windows[i] = entries[i] + entries[i+1] + entries[i+2]
		}
		// count how many sliding window entries are larger than their previous entry
		var sum int
		for i := 0; i < len(sliding_windows)-1; i++ {
			if sliding_windows[i+1] > sliding_windows[i] {
				sum += 1
			}
		}
		return strconv.Itoa(sum)
	}

	return common.Day{ID: "01", Part1: part1, Part2: part2}
}
