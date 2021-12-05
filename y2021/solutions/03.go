package solutions

import (
	"strconv"

	"github.com/codebycaleb/advent-go/common"
	"github.com/codebycaleb/advent-go/common/utils"
)

func reducer(numbers []string, bit int, most_common bool) string {
	if len(numbers) == 1 {
		return numbers[0]
	}

	list_1s, list_0s := []string{}, []string{}
	for _, number := range numbers {
		if number[bit] == '1' {
			list_1s = append(list_1s, number)
		} else {
			list_0s = append(list_0s, number)
		}
	}

	if most_common {
		if len(list_1s) >= len(list_0s) {
			return reducer(list_1s, bit+1, most_common)
		}
		return reducer(list_0s, bit+1, most_common)
	} else {
		if len(list_1s) < len(list_0s) {
			return reducer(list_1s, bit+1, most_common)
		}
		return reducer(list_0s, bit+1, most_common)
	}
}

func Day03() common.Day {

	part1 := func(input string) string {
		gamma_rate, epsilon_rate := 0, 0

		// parse input to lines
		numbers := utils.Lines(input)
		line_length := len(numbers[0])
		bit_counts := make([]int, line_length)

		// record bit counts for each line
		for _, number := range numbers {
			for i, bit := range number {
				bit_counts[i] += int(bit-'0')*2 - 1
			}
		}

		// update the gamma rate for each bit based on the average
		for _, bit_count := range bit_counts {
			gamma_rate <<= 1
			if bit_count > 0 {
				gamma_rate |= 1
			}
		}
		// epsilon rate is the inverse of the gamma rate
		epsilon_rate = ((1 << line_length) - 1) ^ gamma_rate

		return strconv.Itoa(gamma_rate * epsilon_rate)
	}

	part2 := func(input string) string {
		oxygen_generator_rating_string, co2_scrubber_rating_string := "", ""

		// parse input to lines
		numbers := utils.Lines(input)

		oxygen_generator_rating_string = reducer(numbers, 0, true)
		co2_scrubber_rating_string = reducer(numbers, 0, false)

		oxygen_generator_rating, _ := strconv.ParseInt(oxygen_generator_rating_string, 2, 0)
		co2_scrubber_rating, _ := strconv.ParseInt(co2_scrubber_rating_string, 2, 0)

		return strconv.FormatInt(oxygen_generator_rating*co2_scrubber_rating, 10)
	}

	return common.Day{ID: "03", Part1: part1, Part2: part2}
}
