package solutions

import (
	"strconv"
	"strings"

	"github.com/codebycaleb/advent-go/common"
)

func Day06() common.Day {

	type Form struct {
		count     int
		responses map[rune]int
	}

	// split the input into [][]string
	parseInput := func(input string) [][]string {
		groups := strings.Split(input, "\n\n")
		parsed := make([][]string, len(groups))
		for i, group := range groups {
			parsed[i] = strings.Split(group, "\n")
		}
		return parsed
	}

	// map the groups to forms
	mapGroups := func(groups [][]string) []Form {
		forms := make([]Form, len(groups))
		for i, group := range groups {
			forms[i] = Form{
				count:     len(group),
				responses: make(map[rune]int),
			}
			for _, response := range group {
				for _, char := range response {
					forms[i].responses[char]++
				}
			}
		}
		return forms
	}

	part1 := func(input string) string {
		groups := parseInput(input)
		forms := mapGroups(groups)
		count := 0
		for _, form := range forms {
			count += len(form.responses)
		}
		return strconv.Itoa(count)
	}

	part2 := func(input string) string {
		groups := parseInput(input)
		forms := mapGroups(groups)
		count := 0
		for _, form := range forms {
			for _, form_count := range form.responses {
				if form_count == form.count {
					count++
				}
			}
		}
		return strconv.Itoa(count)
	}

	return common.Day{ID: "06", Part1: part1, Part2: part2}
}
