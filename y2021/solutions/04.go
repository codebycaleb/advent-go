package solutions

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/codebycaleb/advent-go/common"
	"github.com/codebycaleb/advent-go/common/utils"
)

func Day04() common.Day {

	type Board struct {
		lines []map[string]struct{}
	}

	parse_boards := func(lines []string) []Board {
		boards := []Board{}
		// chunk the rest of the input in groups of 6 lines (a line break and a 5x5 grid)
		// parse that grid into sets of space-separated numbers and store them in a Board
		for i := 0; i < len(lines); i += 6 {
			board := Board{}
			board.lines = make([]map[string]struct{}, 10)
			for i := 0; i < 10; i++ {
				board.lines[i] = make(map[string]struct{})
			}
			row_strings := lines[i+1 : i+6]
			for i, row_string := range row_strings {
				numbers := utils.Words(row_string)
				for j, number := range numbers {
					board.lines[i][number] = struct{}{}
					board.lines[5+j][number] = struct{}{}
				}
			}
			boards = append(boards, board)
		}
		return boards
	}

	first_winner := func(boards []Board, called_numbers []string) (Board, string, error) {
		// for each called number, remove that number from the rows and columns of each board
		// when a row or column is empty, that board is the first winner
		for _, called_number := range called_numbers {
			for _, board := range boards {
				for _, line := range board.lines {
					delete(line, called_number)
					if len(line) == 0 {
						return board, called_number, nil
					}
				}
			}
		}
		// return error if we didn't find a board
		return Board{}, "", errors.New("no board found")
	}

	last_winner := func(boards []Board, called_numbers []string) (Board, string, error) {
		// store boards in a map so we can remove them easily
		board_map := make(map[int]Board)
		for i, board := range boards {
			board_map[i] = board
		}

		// for each called number, remove that number from the rows and columns of each board
		// when a row or column is empty, that board has won and should be removed from board_map
		// when the last board is removed, return it and the last called number
		for _, called_number := range called_numbers {
			for i, board := range board_map {
				for _, line := range board.lines {
					delete(line, called_number)
					if len(line) == 0 {
						if len(board_map) == 1 {
							return board, called_number, nil
						}
						delete(board_map, i)
						break
					}
				}
			}
		}
		// return error if we didn't find a board
		return Board{}, "", errors.New("no board found")
	}

	calculate_score := func(winning_board Board, called_number string) int {
		// sum all the remaining numbers in the winning board's rows (no need to do columns)
		sum := 0
		for _, row := range winning_board.lines[:5] {
			for number := range row {
				n, _ := strconv.Atoi(number)
				sum += n
			}
		}

		winning_number, _ := strconv.Atoi(called_number)

		return sum * winning_number
	}

	part1 := func(input string) string {
		lines := utils.Lines(input)

		// parse first line (comma separated numbers) into called_numbers
		called_numbers := strings.Split(lines[0], ",")

		// parse the rest of the input into boards
		boards := parse_boards(lines[1:])

		// find the winner (handling any errors)
		winning_board, called_number, err := first_winner(boards, called_numbers)
		if err != nil {
			log.Fatal(err)
			return ""
		}

		// calculate score
		score := calculate_score(winning_board, called_number)

		return strconv.Itoa(score)
	}

	part2 := func(input string) string {
		lines := utils.Lines(input)

		// parse first line (comma separated numbers) into called_numbers
		called_numbers := strings.Split(lines[0], ",")

		// parse the rest of the input into boards
		boards := parse_boards(lines[1:])

		// find the last_winner (handling any errors)
		winning_board, called_number, err := last_winner(boards, called_numbers)
		if err != nil {
			log.Fatal(err)
			return ""
		}

		// calculate score
		score := calculate_score(winning_board, called_number)

		return strconv.Itoa(score)
	}

	return common.Day{ID: "04", Part1: part1, Part2: part2}
}
