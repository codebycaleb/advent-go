package utils

import (
	"fmt"
	"log"
	"os"
)

type Day struct {
	ID    string
	Part1 func(string) string
	Part2 func(string) string
}

func (d Day) Run() {
	// load the file
	location := fmt.Sprintf("inputs/%v.txt", d.ID)
	data, err := os.ReadFile(location)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Day %v Part 1: %v\n", d.ID, d.Part1(string(data)))
	log.Printf("Day %v Part 2: %v\n", d.ID, d.Part2(string(data)))
}
