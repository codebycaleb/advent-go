package solutions

import (
	"fmt"
	"strings"

	"github.com/codebycaleb/advent-go/y2020/utils"
)

func Day2() Day {

	parseLine := func(line string) (int, int, byte, string) {
		// parse from "1-2 a: s" to (1, 2, 'a', "s")
		var (
			a, b int
			c    byte
			str  string
		)
		_, err := fmt.Sscanf(line, "%d-%d %c: %s", &a, &b, &c, &str)
		if err != nil {
			panic(err)
		}
		return a, b, c, str
	}

	checkLine1 := func(line string) bool {
		// parse the line into variables
		a, b, c, str := parseLine(line)
		// count the number of times c appears in str
		count := strings.Count(str, string(c))
		// if the number of times c appears is between a and b inclusive, return true
		return count >= a && count <= b
	}

	checkLine2 := func(line string) bool {
		// parse the line into variables
		a, b, c, str := parseLine(line)
		// check str[a-1] == c
		p1 := str[a-1] == c
		// check str[b-1] == c
		p2 := str[b-1] == c
		// if p1 xor p2, return true
		return p1 != p2
	}

	part1 := func(input string) string {
		// split input into lines
		lines := utils.Lines(input)
		count := 0
		for _, line := range lines {
			if checkLine1(line) {
				count++
			}
		}
		return fmt.Sprintf("%d", count)
	}

	part2 := func(input string) string {
		// split input into lines
		lines := utils.Lines(input)
		count := 0
		for _, line := range lines {
			if checkLine2(line) {
				count++
			}
		}
		return fmt.Sprintf("%d", count)
	}

	return Day{ID: "02", Part1: part1, Part2: part2}
}
