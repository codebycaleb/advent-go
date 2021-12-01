package utils

import (
	"strconv"
	"strings"
)

func Words(input string) []string {
	return strings.Fields(input)
}

func Lines(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func Integers(input string) []int {
	words := Words(input)
	mapped := make([]int, len(words))
	for i, v := range words {
		mapped[i], _ = strconv.Atoi(v)
	}
	return mapped
}
