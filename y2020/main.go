package main

import (
	"flag"
	"log"
	"sort"
	"time"

	"github.com/codebycaleb/advent-go/common"
	"github.com/codebycaleb/advent-go/y2020/solutions"
)

var completedSolutions = map[string]common.Day{
	"01": solutions.Day01(),
	"02": solutions.Day02(),
	"03": solutions.Day03(),
	"04": solutions.Day04(),
	"05": solutions.Day05(),
	"06": solutions.Day06(),
}

func main() {
	// Flags

	// day flag
	var day string
	flag.StringVar(&day, "day", "", "day to run (must be in format DD)")

	// Parse flags
	flag.Parse()

	// Handle day flag
	var daysToRun []common.Day
	if day != "" {
		daysToRun = append(daysToRun, completedSolutions[day])
	} else {
		ids := []string{}
		for id := range completedSolutions {
			ids = append(ids, id)
		}
		sort.Strings(ids)
		for _, id := range ids {
			daysToRun = append(daysToRun, completedSolutions[id])
		}
	}

	start := time.Now()

	log.Println("Advent of Code 2020")
	log.Println("-------------------")

	for _, day := range daysToRun {
		day.Run()
	}

	elapsed := time.Since(start)
	log.Printf("Runtime: %s", elapsed)
}
