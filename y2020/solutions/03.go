package solutions

import (
	"strconv"
	"strings"
)

func Day03() Day {

	// parse input into 2d array of bools where '.' is false and '#' is true
	parse := func(input string) [][]bool {
		lines := strings.Split(input, "\n")
		grid := make([][]bool, len(lines))
		for i, line := range lines {
			grid[i] = make([]bool, len(line))
			for j, c := range line {
				grid[i][j] = c == '#'
			}
		}
		return grid
	}

	// traverse the grid and count the number of trees that are hit by the path of the given slope
	traverse := func(grid [][]bool, slope [2]int) int {
		x, y := 0, 0
		count := 0
		for y < len(grid) {
			if grid[y][x] {
				count++
			}
			// modulo to wrap around the grid since it is infinitely wide
			x = (x + slope[0]) % len(grid[0])
			y += slope[1]
		}
		return count
	}

	part1 := func(input string) string {
		grid := parse(input)
		return strconv.Itoa(traverse(grid, [2]int{3, 1}))
	}

	part2 := func(input string) string {
		grid := parse(input)
		iterations := [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
		product := 1
		for _, i := range iterations {
			product *= traverse(grid, i)
		}
		return strconv.Itoa(product)
	}

	return Day{ID: "03", Part1: part1, Part2: part2}
}
