package solutions

import (
	"strconv"
	"strings"

	"github.com/codebycaleb/advent-go/common"
	"github.com/codebycaleb/advent-go/common/utils"
)

func Day02() common.Day {

	// Instructions are given in the form of one of:
	// forward X | down X | up X
	// where X is the distance to move horizontally or vertically
	type Instruction struct {
		direction string
		distance  int
	}

	part1 := func(input string) string {
		// parse entries into slice of Instruction
		instructions := make([]Instruction, 0)
		entries := utils.Lines(input)
		for _, entry := range entries {
			tokens := strings.Split(entry, " ")
			direction := tokens[0]
			distance, _ := strconv.Atoi(tokens[1])
			instructions = append(instructions, Instruction{direction, distance})
		}

		// starting at (0,0), move based on each instruction
		// "forward" increases x
		// "down" increases y
		// "up" decreases y
		x, y := 0, 0
		for _, instruction := range instructions {
			switch instruction.direction {
			case "forward":
				x += instruction.distance
			case "up":
				y -= instruction.distance
			case "down":
				y += instruction.distance
			}
		}

		// return x * y as a string
		return strconv.Itoa(x * y)
	}

	part2 := func(input string) string {
		// parse entries into slice of Instruction
		instructions := make([]Instruction, 0)
		entries := utils.Lines(input)
		for _, entry := range entries {
			tokens := strings.Split(entry, " ")
			direction := tokens[0]
			distance, _ := strconv.Atoi(tokens[1])
			instructions = append(instructions, Instruction{direction, distance})
		}

		// starting at (0,0) with aim of 0, move based on each instruction
		// "down" increases aim
		// "up" decreases aim
		// "forward" does two things:
		// 		1. increases x by distance
		// 		2. increases y by aim * distance
		x, y, aim := 0, 0, 0
		for _, instruction := range instructions {
			switch instruction.direction {
			case "down":
				aim += instruction.distance
			case "up":
				aim -= instruction.distance
			case "forward":
				x += instruction.distance
				y += aim * instruction.distance
			}
		}

		// return x * y as a string
		return strconv.Itoa(x * y)
	}

	return common.Day{ID: "02", Part1: part1, Part2: part2}
}
